package main

import (
	"app/configs"
	"app/routes"
)

func main() {
	configs.InitDB()
	e := routes.New()
	e.Start(":8000")
}
