package main

import (
	"errors"
	"fmt"
)

func main() {
	//strings
	// Long variable
	var variableName1 string = "hello word"
	//short variable
	variableName2 := "hello world"

	fmt.Println(variableName1)
	fmt.Println(variableName2)

	// primitive: boolean, INT, FLOTAL, string

	//boolean
	boolVar := false
	//%t = typenya apa %v\n  = Valuesnya apa
	fmt.Printf("Type: %T Value: %v\n", boolVar, boolVar)

	// int
	intVar := int(10)
	intVar1 := int32(11)
	intVar2 := int64(12)
	fmt.Printf("Type: %T Value: %v\n", intVar, intVar)
	fmt.Printf("Type: %T Value: %v\n", intVar1, intVar1)
	fmt.Printf("Type: %T Value: %v\n", intVar2, intVar2)

	//float
	floatVar := float32(10.5)
	floatVar1 := float64(11.5)
	fmt.Printf("Type: %T Value: %v\n", floatVar, floatVar)
	fmt.Printf("Type: %T Value: %v\n", floatVar1, floatVar1)

	//bytes alias unit 8 (integer 8)
	bytesVar := byte(255)
	fmt.Printf("Type: %T Value: %v\n", bytesVar, bytesVar)

	bytesVar1 := []byte("Hekakbdabd")
	fmt.Printf("Type: %T Value: %v\n", bytesVar1, bytesVar1)


	//rune alias int32
	runeVar :='💯'
	fmt.Printf("Type: %T Value: %v\n", runeVar, runeVar)

	//complex diperlukan untuk data imajiner
	complexVar := 3 + 4i
	fmt.Printf("Type: %T Value: %v\n", complexVar, complexVar)

	//eror
	errVar := errors.New("errdsadsadw")
	fmt.Printf("Type: %T Value: %v\n",errVar,errVar)

	//interface digunakan sebagai penentu metode dan penyimpanan data kosong
	var myInterfacevar interface{}
	myInterfacevar  = 5
	fmt.Printf("Type: %T Value: %v\n",myInterfacevar,myInterfacevar)
	myInterfacevar  = "safastasfdcs"
	fmt.Printf("Type: %T Value: %v\n",myInterfacevar,myInterfacevar)
}
type MethodList interface {
	MyFunction()
	MyFunction2(int)int
}