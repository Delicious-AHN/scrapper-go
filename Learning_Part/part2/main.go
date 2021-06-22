package main

import (
	"fmt"

	"github.io/Delicious-Ahn/nomad-go/Learning_Part/part2/mydict"
)

func main() {
	dictionary := mydict.Dict{"first": "First word"}
	dictionary["Hello"] = "hello"

	def, err := dictionary.Search("Hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(def)
	}
	def, err = dictionary.Search("Fuck")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(def)
	}

	word := "hello"
	def = "Greeting"
	err = dictionary.Add(word, def)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(dictionary)
	hello, _ := dictionary.Search(word)
	fmt.Println(hello)

	word1 := "hello"
	dictionary.Add(word, "First")
	dictionary.Update(word, "Second")
	if err != nil {
		fmt.Println(err)
	}
	word2, _ := dictionary.Search(word1)
	fmt.Println(word2)
}
