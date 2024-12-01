package aoc

type Stack struct {
	items[]interface{}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	top := s.items[len(s.items) - 1]
	s.items = s.items[:len(s.items)-1]
	return top
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

