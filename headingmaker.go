package main

import (
	"math/rand"
	"strings"
	"time"
)

type HeadingMaker struct {
	MainHeading    []string
	PageHeading    []string
	FontRandomiser func() string
}

func getRandomFontClass() string {
	classList := []string{"class-one", "class-two", "class-three", "class-four", "class-five"}
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(5)
	return classList[index]
}

var HeaderH1 = strings.Split("Jamie Carter", "")
var AboutH1 = strings.Split("About Me", "")
var ProjectsH1 = strings.Split("Selected Projects", "")
var StackH1 = strings.Split("Tech Stack", "")
