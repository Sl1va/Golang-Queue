package Queue

type Queue []int

func (q *Queue) Peek() int {
	return (*q)[0]
}

func (q *Queue) Poll() int {
	value := (*q)[0]
	*q = (*q)[1:]
	return value
}

func (q *Queue) Add(val int){
	*q = append(*q, val)
}