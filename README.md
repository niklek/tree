# tree

`tree` is a cli utility written in Go to list contents of a directory in a tree-like format.
The idea behind was to implement the utility without using recursion calls.
It uses a stack and DFS algorithm to traverse the directory structure.

## Usage

```
./tree --help
Usage of ./tree:
  -a    List all files
  -d    List directories only
  -s    Print the size for each line
```

### Example
Prints the structure of `testdata` directory.

```
./tree -a testdata
└──testdata
    ├──pkg
    |   └──file.txt
    ├──lib
    |   └──godoc
    |       ├──style.css
    |       ├──playground.js
    |       ├──jquery.js
    |       └──images
    |           └──go-logo-blue.svg
    ├──index.html
    └──doc
        └──file.txt
```

### Example with sizes
Very basic for now.
```
./tree -a -s testdata
└──testdata (192 bytes)
    ├──pkg (96 bytes)
    |   └──file.txt (14 bytes)
    ├──lib (96 bytes)
    |   └──godoc (192 bytes)
    |       ├──style.css (161 bytes)
    |       ├──playground.js (158 bytes)
    |       ├──jquery.js (52 bytes)
    |       └──images (96 bytes)
    |           └──go-logo-blue.svg (1472 bytes)
    ├──index.html (131 bytes)
    └──doc (96 bytes)
        └──file.txt (14 bytes)
```
