
package main

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:  %s <public key>\n", os.Args[0])
	}
	
	//filename = os.Args[1]
	
	
}
