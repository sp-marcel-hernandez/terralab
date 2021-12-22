package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("wololo")
	for {
		fmt.Println(os.Getenv("SP_CONVERT"))
		time.Sleep(3 * time.Second)
	}
}
