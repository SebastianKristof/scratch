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

	done := false
	for !done {
		fmt.Print(thisElement.char, "")

		done = thisElement.next == nil
		thisElement = thisElement.next
	}

	fmt.Println()
}
