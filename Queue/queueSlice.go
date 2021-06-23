package main
import "fmt"

// Queue represents a queue
type Queue struct {
	items []int
}

// Enqueue adds integer at the back
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue removes the integer at the front of the queue
// and returns the removed integer
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}
func (q *Queue) dequeue() []int {
	for i := 0; i < len(q.items); i++ {

		if (q.items[i]!=0){
			q.items[i]=0
			break
		}
	}
	return q.items
}


func main() {
	myQueue := Queue{}
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	myQueue.Enqueue(400)
	myQueue.Enqueue(500)
	
	fmt.Println(myQueue)
	myQueue.Dequeue() //remove the index
	myQueue.dequeue() //remove the number 
	fmt.Println(myQueue)

}