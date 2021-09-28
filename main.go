package main

import (
	"fmt"
	"os"
)

func main() {

	sslmode := os.Getenv("ServerCode")
	fmt.Println(sslmode)

}
