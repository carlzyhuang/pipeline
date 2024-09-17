package serial

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// Serializer 的默认实现
// struct 结构需要支持 json 格式
type DefaultSerializer[T comparable] struct {
}

func (a *DefaultSerializer[T]) Marshal(rawID T) ([]byte, error) {
	var raw any = rawID

	v := reflect.ValueOf(rawID)
	switch v.Kind() {
	case reflect.Uint8:
		myID := raw.(uint8)
		return []byte{byte(myID)}, nil
	case reflect.Uint16:
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, raw.(uint16))
		return buf, nil
	case reflect.Uint32:
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, raw.(uint32))
		return buf, nil
	case reflect.Uint64:
		serial := &Uint64Serializer{}
		return serial.Marshal(raw.(uint64))
		// buf := make([]byte, 8)
		// binary.BigEndian.PutUint64(buf, raw.(uint64))
		// return buf, nil
	case reflect.Uint:
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(raw.(uint)))
		return buf, nil
	case reflect.Int8:
		myID := raw.(int8)
		return []byte{byte(myID)}, nil
	case reflect.Int16:
		myID := raw.(int16)
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, uint16(myID))
		return buf, nil
	case reflect.Int32:
		myID := raw.(int32)
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, uint32(myID))
		return buf, nil
	case reflect.Int64:
		myID := raw.(int64)
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(myID))
		return buf, nil
	case reflect.Int:
		myID := raw.(int)
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(myID))
		return buf, nil
	case reflect.String:
		return []byte(raw.(string)), nil
	case reflect.Struct:
		return json.Marshal(rawID)
	}

	return nil, fmt.Errorf("actor id %v type %v not implement", rawID, reflect.TypeOf(rawID))
}

func (a *DefaultSerializer[T]) Unmarshal(data []byte) (T, error) {
	var raw T

	v := reflect.ValueOf(raw)
	switch v.Kind() {
	case reflect.Int8:
		var id uint8
		reader := bytes.NewReader(data)
		err := binary.Read(reader, binary.BigEndian, &id)
		if err != nil {
			return raw, err
		}
		var rawID any = int8(id)
		return rawID.(T), nil
	case reflect.Int16:
		var id uint16
		reader := bytes.NewReader(data)
		err := binary.Read(reader, binary.BigEndian, &id)
		if err != nil {
			return raw, err
		}
		var rawID any = int16(id)
		return rawID.(T), nil
	case reflect.Int32:
		var id uint32
		reader := bytes.NewReader(data)
		err := binary.Read(reader, binary.BigEndian, &id)
		if err != nil {
			return raw, err
		}
		var rawID any = int32(id)
		return rawID.(T), nil
	case reflect.Int64:
		var id uint64
		reader := bytes.NewReader(data)
		err := binary.Read(reader, binary.BigEndian, &id)
		if err != nil {
			return raw, err
		}
		var rawID any = int64(id)
		return rawID.(T), nil
	case reflect.Int:
		var id uint64
		reader := bytes.NewReader(data)
		err := binary.Read(reader, binary.BigEndian, &id)
		if err != nil {
			return raw, err
		}
		var rawID any = int(id)
		return rawID.(T), nil
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		var id T
		reader := bytes.NewReader(data)
		err := binary.Read(reader, binary.BigEndian, &id)
		if err != nil {
			return raw, err
		}
		return id, nil
	case reflect.Uint64:
		serial := &Uint64Serializer{}
		id, err := serial.Unmarshal(data)
		if err != nil {
			return raw, err
		}
		var rawID any = uint64(id)
		return rawID.(T), nil
	case reflect.Uint:
		var id uint64
		reader := bytes.NewReader(data)
		err := binary.Read(reader, binary.BigEndian, &id)
		if err != nil {
			return raw, err
		}
		var rawID any = uint(id)
		return rawID.(T), nil
	case reflect.String:
		var rawID any = string(data)
		return rawID.(T), nil
	case reflect.Struct:
		err := json.Unmarshal(data, &raw)
		if err != nil {
			return raw, err
		}
		return raw, nil
	default:
		return raw, fmt.Errorf("actor id unmarshal fail, type  %v not implement", v.Kind())
	}
}

type ByteSerializer struct {
}

func (s *ByteSerializer) Marshal(id []byte) ([]byte, error) {
	return id, nil
}

func (s *ByteSerializer) Unmarshal(data []byte) ([]byte, error) {
	return data, nil
}

type Uint64Serializer struct {
}

// func (s *Uint64Serializer) Marshal(id uint64) ([]byte, error) {
// 	buf := make([]byte, 8)
// 	binary.BigEndian.PutUint64(buf, id)
// 	return buf, nil
// }

// func (s *Uint64Serializer) Unmarshal(data []byte) (uint64, error) {
// 	return binary.BigEndian.Uint64(data), nil
// }

func (s *Uint64Serializer) Marshal(id uint64) ([]byte, error) {
	return []byte(strconv.FormatUint(id, 10)), nil
}

func (s *Uint64Serializer) Unmarshal(data []byte) (uint64, error) {
	id, err := strconv.ParseUint(string(data), 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}
