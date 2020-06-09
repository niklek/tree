package main

import (
	"bytes"
	"testing"
)

const resultTestdataAll = `└──testdata (192 bytes)
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
`

func TestTestdataAll(t *testing.T) {
	out := new(bytes.Buffer)
	err := listPath(out, "testdata", true, false, true)
	if err != nil {
		t.Errorf("unexpected error:%s", err)
	}
	result := out.String()
	if result != resultTestdataAll {
		t.Errorf("bad result\nexpected:\n%v\ngot:\n%v\n", result, resultTestdataAll)
	}
}
