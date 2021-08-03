package main

import (
	"math/rand"
	"time"
)

func getRandomFontClass() string {
	classList := []string{"class-one", "class-two", "class-three", "class-four", "class-five"}

	rand.Seed(time.Now().UnixNano())

	index := rand.Intn(5)

	// fmt.Println(classList[index])
	return classList[index]
}
