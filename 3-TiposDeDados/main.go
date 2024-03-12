package main

import (
	"errors"
	"fmt"
)

func main() {
	// NUMEROS INTEIROS
	// int, int8, int16, int32, int64
	var negativeNumber int = -10000
	var positiveNumber int = 10000
	fmt.Println(negativeNumber)
	fmt.Println(positiveNumber)

	// uint, uint8, uint16, uint32, uint64 (unsigned int, sempre numeros positivos)
	var alwaysPositive uint = 10000 
	fmt.Println(alwaysPositive)

	//rune === int32
	var runeNumber rune = 123456
	fmt.Println(runeNumber)

	//byte === uint8
	var byteNumber byte = 123
	fmt.Println(byteNumber)

	// NUMEROS REAIS
	// float32, float64
	var float32Number float32 = 123.45
	var float64Number float64 = 12300000000.45
	fmt.Println(float32Number)
	fmt.Println(float64Number)


	// STRINGS
	var str string = "hello world"
	fmt.Println(str)

	//BOOLEAN
	var booleano bool = true
	fmt.Println(booleano)

	//ERROR
	var erro error
	fmt.Println(erro)

	var erroPersonalizado error = errors.New("Erro interno")
	fmt.Println(erroPersonalizado)
}