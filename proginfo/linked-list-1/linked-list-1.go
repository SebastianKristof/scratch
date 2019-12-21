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

	// read runes from input into the linked list, such that the elements are added from the front
	var thisElement, nextElement *listElement
	for s.Scan() {
		fmt.Print(s.Text())
		thisElement = new(listElement)
		thisElement.char = s.Text()
		// implied for 1st element: thisElement.next = nil
		if nextElement != nil {
			thisElement.next = nextElement
		}
		nextElement = thisElement
	}

	// print out the resultant linked list
	for thisElement != nil {
		fmt.Print(thisElement.char)
		thisElement = thisElement.next
	}

	fmt.Println()
}
