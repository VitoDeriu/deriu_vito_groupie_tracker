package main

import (
	"GroupieTrackerJJBA/router"
	"GroupieTrackerJJBA/temps"
)

func main() {
	temps.InitTemplate()
	router.InitServer()
}
