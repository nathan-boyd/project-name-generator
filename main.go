package main

import (
	"fmt"
	"math/rand"
	"time"

	ng "github.com/docker/docker/pkg/namesgenerator"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(ng.GetRandomName(0))
}
