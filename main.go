package main
import "fmt"
//Creating a Queue Structure

type Queue struct {
	front, rear, size, i 	int
	capacity             	int
	slice[]  				int	
}

//function to create a Queue

func newQueue (capacity int)*Queue{
	queue 			:= new(Queue)
	queue.capacity 	= capacity
	queue.front 	= 0
	queue.size 		= 0
	queue.rear 		= capacity - 1
	queue.slice = make([]int, capacity)
	fmt.Println(queue.slice)
	return queue	
}

func isFull(queue *Queue) bool{
	return(queue.size==queue.capacity)
}

func isEmpty(queue *Queue) bool{
	return(queue.size==0)
}

func enqueue(queue *Queue, item int){
	if(isFull(queue)){
		return;
	}else{
		queue.rear=(queue.rear+1)%queue.capacity
		queue.slice[queue.rear]=item
		queue.size=queue.size+1
		fmt.Println(item," Foi enfileirado")
		fmt.Println(queue.slice)
	}
}


func main() {
var queue *Queue = newQueue(10)
enqueue(queue, 1)

}