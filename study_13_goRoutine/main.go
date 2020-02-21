package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	people := [2]string{"paul", "Abigail"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println("waiting for", i)
		fmt.Println(<-c)
	}
}

func isSexy(person string, c chan string) {
	// time.Sleep(time.Second * 10)
	c <- person + " is sexy"
}
