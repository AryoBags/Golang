package main

import "fmt"

func print_hello(){
	fmt.Printf("Hello, World!")
}
func main(){

	//defer
	defer print_hello()

	
	// if else if , else



	age := 25
	if age <10{
		fmt.Printf("Not old enough\n")
	}else if age > 60{
		fmt.Printf("too Old\n")
	}else {
		fmt.Printf("Welcome\n")
	}

	// Switch

	// grade := "h"
	// switch grade{
	// 	case "a":
	// 		fmt.Printf("Nice\n")
	// 	case "d", "b", "c":
	// 		fmt.Printf("NT\n")
	// 		fmt.Printf("GG\n")
	// 		fmt.Printf("Block\n")
	// 	case "g":
	// 		fmt.Printf("hala lel lee\n")
	// 		fallthrough
	// default:
	// 	fmt.Printf("cupu ")
	// }


	// looping\

	// for i := 0; i <5; i++{
	// 	fmt.Printf("looping %d\n", i)
	// }

	// While Looping

	// j := 0 
	// for j < 5 {
	// 	fmt.Printf("looping %d\n", j)
	// 	j+=1
	// }

	// infinity, break , cont  Loop
	
	// i := 0 
	// for {
	// 	i += 1
	// 	if i == 2{
	// 		continue
	// 	}
	// 	fmt.Printf("Looping %d\n", i)
		
	// 	if i == 5 {
	// 		break
	// 	}
	// }


	//range loop litelatur
	
	// number := []int64{1,2,3,4,5}

	// for i, v := range number{
	// 	fmt.Printf("index %d, value %d\n", i, v)
	// }

}