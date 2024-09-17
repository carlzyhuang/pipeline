package jobs

import "container/list"

// Queue 是一个基于 container/list 的队列结构
type Queue struct {
	list *list.List
}

// NewQueue 创建一个新的队列
func NewQueue() *Queue {
	return &Queue{list: list.New()}
}

// Enqueue 在队尾添加一个元素
func (q *Queue) Enqueue(item interface{}) {
	q.list.PushBack(item)
}

// Dequeue 从队首移除一个元素并返回它
func (q *Queue) Dequeue() interface{} {
	if q.list.Len() == 0 {
		return nil
	}
	element := q.list.Front()
	q.list.Remove(element)
	return element.Value
}

// IsEmpty 检查队列是否为空
func (q *Queue) IsEmpty() bool {
	return q.list.Len() == 0
}

// Size 返回队列的大小
func (q *Queue) Size() int {
	return q.list.Len()
}
