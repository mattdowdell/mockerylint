package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/mattdowdell/mockerylint"
)

func main() {
	singlechecker.Main(mockerylint.New())
}
