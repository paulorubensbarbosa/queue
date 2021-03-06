package queue
import "fmt"
import "math"
//Creating a Queue Structure

type Queue struct {
	front, rear, size, i 	int
	capacity             	int
	slice[]  				int	
}

//function to create a Queue

func NewQueue (capacity int)*Queue{
	queue 			:= new(Queue)
	queue.capacity 	= capacity
	queue.front 	= 0
	queue.size 		= 0
	queue.rear 		= capacity - 1
	queue.slice = make([]int, capacity)
	return queue	
}
//functions for verification
func isFull(queue *Queue) bool{
	return(queue.size==queue.capacity)
}

func isEmpty(queue *Queue) bool{
	return(queue.size==0)
}
//Function to enqueue, this add 1 item int tipe, add + 1 to queue.rear and add+ 1 to queue size
func Enqueue(queue *Queue, item int){
	if(isFull(queue)){
		return;
	}else{
		queue.rear=(queue.rear+1)%queue.capacity
		queue.slice[queue.rear]=item
		queue.size=queue.size+1
		fmt.Println("item: ", item," enqueued to queue")
	}
}
//Function to dequeue, this remove 1 item int tipe(the first of the queue), add + 1 to the queue front and remove -1 to de queue size
func Dequeue(queue *Queue) int{
	if(isEmpty(queue)){
		return math.MinInt16
	}else{
		var item int = queue.slice[queue.front]
		queue.front = (queue.front + 1)%queue.capacity
		queue.size=queue.size-1
		return item
	}
}
// functions for show the intem at the front and at the rear of the queue
func GetFront(queue *Queue)int{
	if(isEmpty(queue)){
		return math.MinInt16
	}else{
		return queue.slice[queue.front]
	}
}

func GetRear(queue *Queue)int{
	if(isEmpty(queue)){
		return math.MinInt16
	}else{
		return queue.slice[queue.rear]
	}
}

func GetSize(queue *Queue)int{
	return queue.size
}

//testing all functions
/*func main() {
var queue *Queue = NewQueue(10)
Enqueue(queue, 1)
Enqueue(queue, 2)
Enqueue(queue, 3)

fmt.Println ("item: ", Dequeue(queue), " Dequeued from queue" )
fmt.Println ("item: ", Dequeue(queue), " Dequeued from queue" )
fmt.Println ("item: ", Dequeue(queue), " Dequeued from queue" )

fmt.Println("Front item is: ", GetFront(queue))
fmt.Println("Rear item is: ", GetRear(queue))

fmt.Println(GetSize(queue))
}
*/