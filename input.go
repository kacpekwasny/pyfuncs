package pyfuncs

import (
	"bufio"
	"fmt"
	"os"
)

// Input is a python input function
func Input(prompt string) string {
	fmt.Printf(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
