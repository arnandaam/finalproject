package main

import (
	"finalMygram/app"
	"finalMygram/router"
)

func main() {
	app.StartDB()
	r := router.StartApp()
	r.Run()
}
