package main

import "fmt"

func menu() {
	fmt.Println("🕸 Pick what kind of job would you like to perform:")
	options := map[int]string{
		1: "Record Book",
		2: "Delete Book"}

	for option := range options {
		fmt.Printf(" OPTION: [%d] ▶️ %s\n", option, options[option])
	}
}

func userOption() {
	var userJobPicked string
	fmt.Print("[OPTION]:")
	fmt.Scanln(&userJobPicked)

	//options call functions
}

func main() {
	fmt.Println("Welcome to GutenbergDB")
	menu()
	userOption()
}
