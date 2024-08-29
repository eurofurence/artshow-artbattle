package main

import (
	"github.com/eurofurence/artshow-artbattle/internal/application/app"
	"os"
)

func main() {
	os.Exit(app.New().Run())
}
