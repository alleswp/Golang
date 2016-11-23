package utilities

import (
	"fmt"
"strings"
)
func PrintHeader(str string) {
	fmt.Println("\n\n----------------------------------------------------------------")
	fmt.Println("\t\t", strings.ToUpper(str))
	fmt.Println("----------------------------------------------------------------")
}
