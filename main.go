package main

import "fmt"

func main() {
	var input Word
	fmt.Print("Input a word: ")
	fmt.Scan(&input.word)

	polindrome, err := input.polindromeDeterminer()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(polindrome)
}

type Word struct {
	word string
}

func (w Word) polindromeDeterminer() (string, error) {
	//error handler for only one character
	if len(w.word) <= 1 {
		return "", fmt.Errorf("No word has been enter")
	}

	//reversing input
	output := []rune(w.word)
	for i, j := 0, len(output)-1; i < j; i, j = i+1, j-1 {
		output[i], output[j] = output[j], output[i]
	}

	//condition for polindrome or not
	if string(output) == w.word {
		return string("Polindrome: " + string(output)), nil
	} else {
		return string("Not Polindrome: " + string(output)), nil
	}
}
