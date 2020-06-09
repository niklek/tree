// Stack for traversing a directory
package main

// item represent a file or directory in a stack
type item struct {
	path   string
	prefix string
	isLast bool
}

// a stack to keep traversed items
type Stack struct {
	items []*item
}

func (s *Stack) Push(it *item) {
	s.items = append(s.items, it)
}

func (s *Stack) Pop() *item {
	if s.IsEmpty() {
		return &item{}
	}

	it := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return it
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}
