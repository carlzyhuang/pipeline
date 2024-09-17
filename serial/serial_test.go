package serial

import (
	"math"
	"testing"
)

func TestSerialInt32(t *testing.T) {
	cases := []struct {
		Name    string
		HashKey int32
	}{
		{
			"test 1",
			math.MinInt32,
		},
		{
			"test 2",
			math.MaxInt32,
		},
		{
			"test 3",
			0,
		},
	}
	serial := &DefaultSerializer[int32]{}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			id, err := serial.Marshal(c.HashKey)
			if err != nil {
				t.Fatalf("%s: marshal got err %v", c.Name, err)
			}

			orignal, err := serial.Unmarshal(id)
			if err != nil {
				t.Fatalf("%s: unmarshal got err %v", c.Name, err)
			}

			if orignal != c.HashKey {
				t.Fatalf("%s: expected %v, got %v", c.Name, c.HashKey, orignal)
			}
		})
	}
}

func TestSerialUint8(t *testing.T) {
	cases := []struct {
		Name    string
		HashKey uint8
	}{
		{
			"test 1",
			math.MaxUint8,
		},
		{
			"test 3",
			0,
		},
	}
	serial := &DefaultSerializer[uint8]{}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			id, err := serial.Marshal(c.HashKey)
			if err != nil {
				t.Fatalf("%s:  got err %v", c.Name, err)
			}

			orignal, err := serial.Unmarshal(id)
			if err != nil {
				t.Fatalf("%s:  got err %v", c.Name, err)
			}

			if orignal != c.HashKey {
				t.Fatalf("%s: expected %v, got %v", c.Name, c.HashKey, orignal)
			}
		})
	}
}

func TestSerialUint32(t *testing.T) {
	cases := []struct {
		Name    string
		HashKey uint32
	}{
		{
			"test 1",
			math.MaxUint32,
		},
		{
			"test 3",
			0,
		},
	}
	serial := &DefaultSerializer[uint32]{}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			id, err := serial.Marshal(c.HashKey)
			if err != nil {
				t.Fatalf("%s:  got err %v", c.Name, err)
			}

			orignal, err := serial.Unmarshal(id)
			if err != nil {
				t.Fatalf("%s:  got err %v", c.Name, err)
			}

			if orignal != c.HashKey {
				t.Fatalf("%s: expected %v, got %v", c.Name, c.HashKey, orignal)
			}
		})
	}
}

func TestSerialUint64(t *testing.T) {
	cases := []struct {
		Name    string
		HashKey uint64
	}{
		{
			"test 1",
			math.MaxUint64,
		},
		{
			"test 3",
			0,
		},
	}
	serial := &DefaultSerializer[uint64]{}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			id, err := serial.Marshal(c.HashKey)
			if err != nil {
				t.Fatalf("%s:  got err %v", c.Name, err)
			}

			orignal, err := serial.Unmarshal(id)
			if err != nil {
				t.Fatalf("%s:  got err %v", c.Name, err)
			}

			if orignal != c.HashKey {
				t.Fatalf("%s: expected %v, got %v", c.Name, c.HashKey, orignal)
			}
		})
	}
}

func TestSerialUint(t *testing.T) {
	cases := []struct {
		Name    string
		HashKey uint
	}{
		{
			"test 1",
			math.MaxUint,
		},
		{
			"test 3",
			0,
		},
	}
	serial := &DefaultSerializer[uint]{}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			id, err := serial.Marshal(c.HashKey)
			if err != nil {
				t.Fatalf("%s:  got err %v", c.Name, err)
			}

			orignal, err := serial.Unmarshal(id)
			if err != nil {
				t.Fatalf("%s:  got err %v", c.Name, err)
			}

			if orignal != c.HashKey {
				t.Fatalf("%s: expected %v, got %v", c.Name, c.HashKey, orignal)
			}
		})
	}
}

func TestSerialint8(t *testing.T) {
	cases := []struct {
		Name    string
		HashKey int8
	}{
		{
			"test 1",
			math.MinInt8,
		},
		{
			"test 2",
			math.MaxInt8,
		},
		{
			"test 3",
			0,
		},
	}
	serial := &DefaultSerializer[int8]{}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			id, err := serial.Marshal(c.HashKey)
			if err != nil {
				t.Fatalf("%s:  got err %v", c.Name, err)
			}

			orignal, err := serial.Unmarshal(id)
			if err != nil {
				t.Fatalf("%s:  got err %v", c.Name, err)
			}

			if orignal != c.HashKey {
				t.Fatalf("%s: expected %v, got %v", c.Name, c.HashKey, orignal)
			}
		})
	}
}

func TestSerialint16(t *testing.T) {
	cases := []struct {
		Name    string
		HashKey int16
	}{
		{
			"test 1",
			math.MinInt16,
		},
		{
			"test 2",
			math.MaxInt16,
		},
		{
			"test 3",
			0,
		},
	}
	serial := &DefaultSerializer[int16]{}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			id, err := serial.Marshal(c.HashKey)
			if err != nil {
				t.Fatalf("%s:  got err %v", c.Name, err)
			}

			orignal, err := serial.Unmarshal(id)
			if err != nil {
				t.Fatalf("%s:  got err %v", c.Name, err)
			}

			if orignal != c.HashKey {
				t.Fatalf("%s: expected %v, got %v", c.Name, c.HashKey, orignal)
			}
		})
	}
}

func TestSerialint32(t *testing.T) {
	cases := []struct {
		Name    string
		HashKey int32
	}{
		{
			"test 1",
			math.MinInt32,
		},
		{
			"test 2",
			math.MaxInt32,
		},
		{
			"test 3",
			0,
		},
	}
	serial := &DefaultSerializer[int32]{}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			id, err := serial.Marshal(c.HashKey)
			if err != nil {
				t.Fatalf("%s:  got err %v", c.Name, err)
			}

			orignal, err := serial.Unmarshal(id)
			if err != nil {
				t.Fatalf("%s:  got err %v", c.Name, err)
			}

			if orignal != c.HashKey {
				t.Fatalf("%s: expected %v, got %v", c.Name, c.HashKey, orignal)
			}
		})
	}
}

func TestSerialint64(t *testing.T) {
	cases := []struct {
		Name    string
		HashKey int64
	}{
		{
			"test 1",
			math.MinInt64,
		},
		{
			"test 2",
			math.MaxInt64,
		},
		{
			"test 3",
			0,
		},
	}
	serial := &DefaultSerializer[int64]{}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			id, err := serial.Marshal(c.HashKey)
			if err != nil {
				t.Fatalf("%s:  got err %v", c.Name, err)
			}

			orignal, err := serial.Unmarshal(id)
			if err != nil {
				t.Fatalf("%s:  got err %v", c.Name, err)
			}

			if orignal != c.HashKey {
				t.Fatalf("%s: expected %v, got %v", c.Name, c.HashKey, orignal)
			}
		})
	}
}

func TestSerialString(t *testing.T) {
	cases := []struct {
		Name    string
		HashKey string
	}{
		{
			"test 1",
			"",
		},
		{
			"test 2",
			"aabbcc",
		},
		{
			"test 3",
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		},
	}
	serial := &DefaultSerializer[string]{}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			id, err := serial.Marshal(c.HashKey)
			if err != nil {
				t.Fatalf("%s:  got err %v", c.Name, err)
			}

			orignal, err := serial.Unmarshal(id)
			if err != nil {
				t.Fatalf("%s:  got err %v", c.Name, err)
			}

			if orignal != c.HashKey {
				t.Fatalf("%s: expected %v, got %v", c.Name, c.HashKey, orignal)
			}
		})
	}
}
