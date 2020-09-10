package main

import (
	"../../models"
	"fmt"
)


func main() {
	disc := models.Disc{Sequence: "abcdefghijklmnopqrstuvwxyz"}
	//test1 := test[2]
	fmt.Printf("%d",len(disc.Sequence))
	//fmt.Printf("%x", test1)
}
