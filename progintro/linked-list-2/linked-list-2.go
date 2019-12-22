package main

import (
	"bufio"
	"fmt"
	"os"
)

type listElement struct {
	char string
	next *listElement
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanRunes)

	var this, first, last *listElement

	for s.Scan() {

		this = new(listElement)
		this.char = s.Text()
		this.next = nil // the new last element

		if last != nil { // if previous last exists
			last.next = this // connect previous last element to this
		}

		if first == nil {
			first = this
		} // first will stay unchanged after 1st iteration

		last = this
	}

	this = first
	for this != last {
		fmt.Printf("%s, %p, %p\n", this.char, this, this.next)
		this = this.next
	}

}
