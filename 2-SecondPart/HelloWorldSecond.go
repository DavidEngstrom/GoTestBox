package main

import (
	"bufio"
	"fmt"
	"os"
)

func ScanKeybord(textToShow []string, inputTextToShow []string) []string {

	/*
	   // before the loop
	   output := make([]string, len(input))
	   would be my recommendation because append causes a bunch of needless reallocations and you already know what capacity you need since it's based on the input.

	   The other thing would be:

	   output = append(output, input[index])
	   But like I said, from what I've observed append grows the initial capacity exponentially. This will be base 2 if you haven't specified anything which means you're going to be doing several unneeded reallocations before reaching the desired capacity.
	*/

	scanner := bufio.NewScanner(os.Stdin)
	out := make([]string, len(inputTextToShow))
	//var readStr string

	//SHow the output strings
	for i := 0; i < len(textToShow); i++ {
		fmt.Println(textToShow[i])
	}
	for i := 0; i < len(inputTextToShow); i++ {
		fmt.Print(inputTextToShow[i])
		scanner.Scan()
		out[i] = scanner.Text()
	}
	return out
}

/* MainMenue */
func MainMenue() {
	fmt.Println("1 - Calculator ")
}

func main() {
	var text string
	var ret []string
	for text != "q" { // break the loop if text == "q"

		//scanner.Scan()
		var out = []string{"------------", "Top meny ", "1 - Calculator ", "q - Quit "}
		var in = []string{"Select a number"}
		ret = ScanKeybord(out, in) //scanner.Text()
		text = ret[1]
		if text != "q" {
			//fmt.Println("Your text was: ", text)
		}
	}
	fmt.Println("Exit thanks!")
}
