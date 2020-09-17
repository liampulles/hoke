package main

import (
	"os"

	"github.com/liampulles/hoke/pkg/app"
)

func main() {
	os.Exit(app.Run(os.Args))
}
