package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Por gentileza, entre com o seu nome:")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Oiê, eu sou o(a), %s! Vamos nessa!\n", name)
	name = strings.TrimSpace(name)
}