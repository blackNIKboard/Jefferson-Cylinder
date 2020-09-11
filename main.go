//nolint:gochecknoglobals

package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var (
	action   = flag.Bool("action", false, "action to perform (true to decrypt)")
	cypher   = flag.String("file", "", "filename")
	sequence = flag.String("sequence", "", "sequence to encrypt/decrypt")
	position = flag.Int("position", 0, "position to read encrypted")
)

func main() {
	flag.Parse()

	cyl := new(Cylinder)

	if *action == false {
		cyl.init(len(*sequence))
		res := cyl.encode(*sequence, *position)
		cyl.storeShuffle(*cypher)
		fmt.Println(res)
		fmt.Println()
		freqAnalyze(res)
	}

	if *action == true {
		cyl.read(*cypher)
		cyl.decode(*sequence)
	}
}

type Disc struct {
	Sequence string
}

type Cylinder struct {
	Discs  []Disc
	Height int
}

func (d *Disc) rotate(letter rune) {
	inRune := []rune(d.Sequence)
	for inRune[0] != letter {
		temp := inRune[len(inRune)-1]

		for i := len(inRune) - 1; i > 0; i-- {
			inRune[i] = inRune[i-1]
		}

		inRune[0] = temp
	}
	d.Sequence = string(inRune)
}

func freqAnalyze(str string) {
	results := [len(alphabet)]int{0}

	for i := 0; i < len(str); i++ {
		for j := 0; j < len(alphabet); j++ {
			if str[i] == alphabet[j] {
				results[j]++
			}
		}
	}

	sum := 0

	for i := 0; i < len(alphabet); i++ {
		sum += results[i]
		if results[i] != 0 {
			fmt.Printf("%c = %d\n", alphabet[i], results[i])
		}
	}

	fmt.Printf("\nlength = %d\n", sum)
}

func (c *Cylinder) store(filename string) {
	cypher, err := os.Create(filename)

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer cypher.Close()

	cypher.WriteString(strconv.Itoa(c.Height) + "\n")

	for i := 0; i < c.Height; i++ {
		_, err = cypher.WriteString(c.Discs[i].Sequence + "\n")
		if err != nil {
			fmt.Println("Unable to create file:", err)
			os.Exit(1)
		}
	}
}

func (c *Cylinder) read(filename string) {
	cypher, err := os.Open(filename)

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer cypher.Close()

	buffer := bufio.NewReader(cypher)

	height, _ := buffer.ReadString('\n')
	height = strings.TrimSuffix(height, "\n")
	test, _ := strconv.Atoi(height)
	c.Height = test
	c.Discs = make([]Disc, c.Height)

	for i := 0; i < c.Height; i++ {
		temp, _ := buffer.ReadString('\n')
		c.Discs[i].Sequence = strings.TrimSuffix(temp, "\n")
	}
}

func (c *Cylinder) init(height int) {
	c.Height = height
	c.Discs = make([]Disc, height)

	for i := 0; i < height; i++ {
		c.Discs[i].Sequence = alphabet

		c.Discs[i].shuffle()
	}
}

func (c *Cylinder) print() {
	for i := 0; i < c.Height; i++ {
		inRune := []rune(c.Discs[i].Sequence)
		for j := 0; j < len(inRune); j++ {
			fmt.Print(string(inRune[j]) + " | ")
		}
		fmt.Println()
	}
}

func (c *Cylinder) encode(text string, position int) string {
	inRune := []rune(text)
	result := []rune(text)

	for i := 0; i < len(inRune); i++ {
		c.Discs[i].rotate(inRune[i])
		temp := []rune(c.Discs[i].Sequence)
		result[i] = temp[position]
	}

	return string(result)
}

func (c *Cylinder) storeShuffle(filename string) {
	for i := 0; i < c.Height; i++ {
		c.Discs[i].shuffle()
	}
	c.store(filename)
}

func (c *Cylinder) decode(cypher string) {
	inRune := []rune(cypher)

	for i := 0; i < c.Height; i++ {
		c.Discs[i].rotate(inRune[i])
	}

	c.print()
}

func (d *Disc) shuffle() {
	inRune := []rune(d.Sequence)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for n := len(inRune); n > 0; n-- {
		randIndex := r.Intn(n)
		inRune[n-1], inRune[randIndex] = inRune[randIndex], inRune[n-1]
	}

	d.Sequence = string(inRune)
}
