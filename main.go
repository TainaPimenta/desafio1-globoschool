package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Fprint(os.Stdout, "Por gentileza, entre com o seu nome:")
		name, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		name = strings.TrimSpace(name)
		fmt.Printf("Oiê, eu sou o(a), %s! Vamos nessa!\n", name)

	}

}
