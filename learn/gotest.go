package gotest

import (
	"fmt"
	"time"
)

func gotest() {
	c := make(chan string)
	people := [5]string{"nico", "flynn", "youjin", "Jiun", "dongdong"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println("waiting for a message", i)
		fmt.Println(<-c)
	}
}

func isSexy(person string, channel chan string) {
	time.Sleep(time.Second * 10)
	channel <- person + " is sexy"
}
