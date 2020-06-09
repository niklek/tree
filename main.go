// Cli utility to list contents of a directory in a tree-like format
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Prints graphic symbols for an item
func printBar(last bool) string {
	if last {
		return "└──"
	}
	return "├──"
}

func main() {
	var path, fullPath, prefix string

	path = "." // default path
	if len(os.Args) >= 2 {
		path = os.Args[1]
	}
	prefix = ""

	// Create a stack and push the root item
	var s Stack
	s.Push(&item{path, prefix, true})

	for s.Size() > 0 {
		it := s.Pop()
		if it.path == "" {
			break
		}
		// get FileInfo describing the current item
		fi, err := os.Stat(it.path)
		if err != nil {
			log.Println(err)
			break
		}
		// get child items if the item is a directory
		if fi.IsDir() {
			files, err := ioutil.ReadDir(it.path)
			if err != nil {
				log.Println(err)
				break
			}

			for k, file := range files {
				// build full path for the child item
				fullPath = it.path + string(os.PathSeparator) + file.Name()
				// extend graphic prefix for the child item based on parent
				prefix = it.prefix + "|   "
				if it.isLast {
					prefix = it.prefix + "    "
				}
				// push the child into the stack
				s.Push(&item{
					path:   fullPath,
					prefix: prefix,
					isLast: k == 0, // first child item will be the last one when printing
				})
			}
		}

		// print the current item
		fmt.Printf("%s%s%s\n", it.prefix, printBar(it.isLast), fi.Name())
	}
}
