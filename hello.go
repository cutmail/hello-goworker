package main

import (
	"fmt"

	"github.com/benmanns/goworker"
)

func init() {
	goworker.Register("Hello", helloWorker)
}

func main() {
	if err := goworker.Work(); err != nil {
		fmt.Println("Error:", err)
	}
}

func helloWorker(queue string, args ...interface{}) error {
	fmt.Println("Hello, world!")
	return nil
}
