package main

import "fmt"

type Queue []int

func (q *Queue) Peek() int {
	return (*q)[0]
}

func (q *Queue) Poll() int {
	value := (*q)[0]
	*q = (*q)[1:]
	return value
}

func (q *Queue) Add(val ... int) {
	*q = append(*q, val...)
}

func (q *Queue) Size() int {
	return len(*q)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func main() {
	queue := make(Queue, 0) // create new Queue
	fmt.Println("This is empty Queue ", queue)
	fmt.Println("		IsEmpty: ", queue.IsEmpty())
	queue.Add(18, 99, 44, 13, 22, -67)

	fmt.Println("\nThis is filled Queue: ", queue)
	fmt.Println("\nLet's check first element without removing:")
	fmt.Println("		first element: ", queue.Peek())
	fmt.Println("		Queue: ", queue)
	fmt.Println("		Queue size: ", queue.Size())
	fmt.Println("		IsEmpty: ", queue.IsEmpty())
	fmt.Println("\nLet's remove  and second elements:")
	fmt.Println("		first element: ", queue.Poll())
	fmt.Println("		second element: ", queue.Poll())
	fmt.Println("		Queue: ", queue)
	fmt.Println("		Queue size: ", queue.Size())
	fmt.Println("		IsEmpty: ", queue.IsEmpty())
}