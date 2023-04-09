package main

import (
	"fmt"
	"path"
)

func SeparatePath(p string) {
	dir, file := path.Split(p)

	fmt.Println("dir:", dir)
	fmt.Println("file:", file)
}
