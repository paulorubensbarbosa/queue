package main

//Creating a Queue Structure

type Queue struct {
	front, rear, size, i 	int
	capacity             	int
	array[i]				int
}

//function to create a Queue

func newQueue(capacity int) {
	queue 			:= new(Queue)
	queue.capacity 	= capacity
	queue.front 	= 0
	queue.size 		= 0
	queue.rear 		= capacity - 1
	//não sei como colocar o capacity dentro dos ([]) do array
	//queue.array	=  
	queue.i			= capacity
}

func main() {
}
