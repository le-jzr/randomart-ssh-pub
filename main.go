package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/zarevucky/randomart"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:  %s <SSH public key>\n", os.Args[0])
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	for data[0] != ' ' {
		data = data[1:]
	}
	data = data[1:]

	for i := 0; i < len(data); i++ {
		if data[i] == ' ' {
			data = data[:i]
			break
		}
	}

	decoded := make([]byte, base64.StdEncoding.DecodedLen(len(data)))

	n, err := base64.StdEncoding.Decode(decoded, data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	decoded = decoded[:n]
	sum := md5.Sum(decoded)

	lines := randomart.OpenSSH(sum[:])

	fmt.Println("+-----------------+")

	for l := 0; l < len(lines); l++ {
		fmt.Print("|", string(lines[l][:]), "|\n")
	}

	fmt.Println("+-----------------+")
}
