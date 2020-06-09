// Cli utility to list contents of a directory in a tree-like format
package main

import (
	"flag"
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

// Checks whether we need to skip the item or not based on flags
func skipItem(fi os.FileInfo, listAll bool, listDir bool) bool {
	if listDir && fi.Mode().IsRegular() {
		return true
	}

	if !listAll {
		name := fi.Name()
		if len(name) > 1 && name[:1] == "." {
			return true
		}
	}

	return false
}

// Returns a string with file/direcrory size in bytes
// When the flag is false will return an empty string
func printSize(fi os.FileInfo, displaySize bool) string {
	if !displaySize {
		return ""
	}

	return fmt.Sprintf(" (%d bytes)", fi.Size())
}

const pathDefault = "."

func main() {
	var (
		listAll, listDir, displaySize bool
		path, fullPath, prefix        string
	)
	flag.BoolVar(&listAll, "a", false, "List all files")
	flag.BoolVar(&listDir, "d", false, "List directories only")
	flag.BoolVar(&displaySize, "s", false, "Print the size for each line")
	flag.Parse()

	// get path or use default
	args := flag.Args()
	path = pathDefault
	if len(args) > 0 {
		path = args[0]
	}

	// Create a stack and push the root item
	var s Stack
	s.Push(&item{path, "", true})

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
		// when only directories
		if skipItem(fi, listAll, listDir) {
			continue
		}
		// get child items if the item is a directory
		if fi.IsDir() {
			files, err := ioutil.ReadDir(it.path)
			if err != nil {
				log.Println(err)
				break
			}

			for k, file := range files {
				// when only directories
				if skipItem(fi, listAll, listDir) {
					continue
				}
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
		fmt.Printf("%s%s%s%s\n", it.prefix, printBar(it.isLast), fi.Name(), printSize(fi, displaySize))
	}
}
