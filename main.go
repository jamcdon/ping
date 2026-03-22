package main

import (
	"github.com/jamcdon/ping/Routes"
)

func main() {
	router := Routes.SetupRouter()
	router.run()
}