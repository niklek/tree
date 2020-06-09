package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type item struct {
	path   string
	prefix string
	isLast bool
}

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

func printBar(last bool) string {
	if last {
		return "└───"
	}
	return "├───"
}

func main() {
	var path, fullPath, prefix string

	path = "." // default
	if len(os.Args) >= 2 {
		path = os.Args[1]
	}
	prefix = ""

	// DFS
	var s Stack
	s.Push(&item{path, prefix, true})

	for s.Size() > 0 {
		it := s.Pop()
		if it.path == "" {
			break
		}

		fi, err := os.Stat(it.path)
		if err != nil {
			log.Println(err)
			break
		}

		if fi.IsDir() {
			files, err := ioutil.ReadDir(it.path)
			if err != nil {
				log.Println(err)
				break
			}

			for k, file := range files {
				fullPath = it.path + string(os.PathSeparator) + file.Name()
				prefix = it.prefix + "|   "
				if it.isLast {
					prefix = it.prefix + "    "
				}
				s.Push(&item{
					path:   fullPath,
					prefix: prefix,
					isLast: k == 0,
				})
			}
		}

		fmt.Printf("%s%s%s\n", it.prefix, printBar(it.isLast), fi.Name())
	}
}

//if file.Mode().IsRegular() {
