package main

import "fmt"

func main() {
	lang := make(map[string]string)

	lang["JS"] = "JavaScript"
	lang["RB"] = "Ruby"
	lang["PY"] = "Python"
	lang["GO"] = "Golang"

	fmt.Println("List of all Language :", lang)
	fmt.Println("JS has :", lang["JS"])

	delete(lang, "RB")

	for key, value := range lang {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}
}
