package main

import (
	router "github.com/alifudin-a/go-ocr/routes"
)

func main() {
	e := router.Init()
	e.Logger.Fatal(e.Start(":7890"))
}
