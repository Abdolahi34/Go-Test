package main

import "fmt"

const var_1 string = "Hello !!!"

func main() {
	fmt.Println(var_1)
	// Declaring Variables
	fmt.Print("\nDeclaring Variables:\n")
	var_2 := 1
	fmt.Printf("var: %v, type: %T\n", var_2, var_2)
	var var_3 int8 = 15
	fmt.Printf("var: %v, type: %T\n", var_3, var_3)
	var var_4 float32 = 1564
	fmt.Print("var: ")
	fmt.Printf("%v, ", var_4)
	fmt.Print("type: ")
	fmt.Printf("%T\n", var_4)
	var_5 := true
	fmt.Println(var_5)
	// Declaring Array
	fmt.Print("\nDeclaring Array:\n")
	var var_6 = [...]string{"Lorem", "ipsum", "is", "placeholder", "text", "commonly", "used", "in", "the", "graphic"}
	fmt.Printf("var: %v\ntype: %T\n", var_6, var_6)
	fmt.Println("length: ", len(var_6))
	fmt.Println("capacity: ", cap(var_6))
	// Declaring Slice
	fmt.Print("\nDeclaring Slice:\n")
	var_7 := []uint16{5, 47, 97, 157, 457, 8957, 14578}
	fmt.Printf("var: %v\ntype: %T\n", var_7, var_7)
	fmt.Println("length: ", len(var_7))
	fmt.Println("capacity: ", cap(var_7))
}
