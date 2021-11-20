package packages

import (
	"fmt"
	"math/rand"
)

// Every go program is made up of packages

// Program start running in package main.

func Main() {
	fmt.Println("Package - My favorite number is", rand.Intn(10))
}
