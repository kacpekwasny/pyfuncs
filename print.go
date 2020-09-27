package pyfuncs

import (
	"fmt"
)

// Print like python print
func Print(valuesIn ...interface{}) {
	fmt.Println(valuesIn...)
}
