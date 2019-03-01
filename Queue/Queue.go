package Queue

type Queue []interface{}

func NewQueue() Queue{
	queue := Queue(make([]interface{}, 0))
	return queue
}

func (q *Queue) Peek() interface{} {
	return (*q)[0]
}

func (q *Queue) Poll() interface{} {
	value := (*q)[0]
	*q = (*q)[1:]
	return value
}

func (q *Queue) Add(val ... interface{}) {
	*q = append(*q, val...)
}

func (q *Queue) Size() int {
	return len(*q)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}