package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("wololo")
	fmt.Println(os.Getenv("SP_CONVERT"))
}
