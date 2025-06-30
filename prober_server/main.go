package main 

import (
	"fmt"
)

func main() {
	res, err := http.Get("")
	if err != nil {
		log.Fatal(err)
	}
}