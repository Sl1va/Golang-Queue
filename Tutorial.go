package main

import (
	. "./Queue"
	"fmt"
)


func main() {
	queue := NewQueue()
	fmt.Println("This is empty Queue ", queue)
	fmt.Println("		IsEmpty: ", queue.IsEmpty())
	queue.Add(18, "FFF", "abcde", 13, "HelloQueue", -67)


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

	fmt.Println("\n//////////////////////////////\n")

	fmt.Println("Let's try to create int Queue")
	intQueue := NewQueue()
	intQueue.Add(10, 55, 67, 99, 32, 33)
	fmt.Println("\nQueue: ", intQueue)
	fmt.Println("Let's try to get sum of all elements")
	sum := 0
	for !intQueue.IsEmpty(){
		sum += intQueue.Poll().(int)
	}
	fmt.Println("		sum = ", sum)
}
