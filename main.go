package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Printf("time:+%v\n", time.Now().Format(time.RFC822))
		time.Sleep(5 * time.Second)
	}

}

func demo() {
	fmt.Println("demo")
}
