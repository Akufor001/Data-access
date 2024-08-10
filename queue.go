package main

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1 // or return an error
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) Len() int {
	return len(q.items)
}

//func queue() {
// Create a new queue
//  queue := &Queue{}

// Enqueue some elements
//queue.Enqueue(1)
//queue.Enqueue(2)
//queue.Enqueue(3)

// Dequeue and display the elements
//for queue.Len() > 0 {
//  fmt.Println(queue.Dequeue())
//}
//}
