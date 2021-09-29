package main

import (
	"fmt"
	"os"
)

func main() {

	sslmode := os.Getenv("ServerCode")
	// SAP_USERNAME SAP_PASSWORD
	fmt.Println(sslmode)

}
