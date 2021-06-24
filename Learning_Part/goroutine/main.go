package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [5]string{"nico", "flynn", "lalal", "japanguy", "dlallt"}
	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println("Waiting for messages")
	// result1 := <-c
	// result2 := <-c
	// result3 := <-c
	// fmt.Println("Received this message 1: ", result1) // waiting a message, if this line get the message, go to next line
	// fmt.Println("Received this message 2: ", result2) // also same
	// fmt.Println("Received this message 3: ", result3)

	for i := 0; i < len(people); i++ {
		fmt.Println("Wating for ", i)
		fmt.Println(<-c)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- person + "is sexy"
}
