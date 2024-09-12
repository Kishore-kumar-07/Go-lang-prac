package main

import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Enter Somrthing:");
	input , _ := reader.ReadString('\n');

	// the 2 are different
	fmt.Println("Somting is ",input);
	fmt.Printf("Something is: %s", input)
}