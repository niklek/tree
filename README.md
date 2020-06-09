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
    ├──doc
    |   └──file.txt
    └──.DS_Store
```

