package main
import "fmt"
//Creating a Queue Structure

type Queue struct {
	front, rear, size, i 	int
	capacity             	int
	slice[]  				int	
}

//function to create a Queue

func newQueue(capacity int) {
	queue 			:= new(Queue)
	queue.capacity 	= capacity
	queue.front 	= 0
	queue.size 		= 0
	queue.rear 		= capacity - 1
	queue.slice = make([]int, capacity)
	
	fmt.Println("item no slice: ", queue.slice[0])
	fmt.Println("cap: ", cap(queue.slice))
}


func main() {
	newQueue(10)
}
