package main

import "fmt"

func main() {
	var variavel string = "variavel1"
	variavel2 := "variavel2"
	fmt.Println(variavel)
	fmt.Println(variavel2)

	var(variavel3 string = "variavel3"
	variavel4 string)

	variavel4 = "variavel4"
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel5", "variavel6"
	fmt.Println(variavel5, variavel6)

	const constante string = "Constante1"
	fmt.Println(constante)


	variavel5, variavel6 = variavel6, variavel5 //mudando a atribuição das variaveis

	fmt.Println(variavel5, variavel6)
}