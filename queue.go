package aoc

type Queue struct {
	items[]interface{}
}

func (q *Queue) Enqueue(item interface{}) {
	s.items = append(q.items, item)
}

func (q *Queue) Dequeue() interface{} {
	if s.IsEmpty() {
		return nil
	}
	bottom := s.items[0]
	s.items = s.items[1:]
	return bottom
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}

