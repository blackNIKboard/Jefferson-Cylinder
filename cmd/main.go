package main

import (
	"fmt"

	"github.com/blackNIKboard/Jefferson-Cylinder/models/cylinder"
)

func main() {
	disc := cylinder.Disc{Sequence: "abcdefghijklmnopqrstuvwxyz"}
	//test1 := test[2]
	fmt.Printf("%d", len(disc.Sequence))
	//fmt.Printf("%x", test1)
}
