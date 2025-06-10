package main

import "url-shortener/app"

func main() {
	app := app.NewApp()
	if err := app.Router.Run(":8080"); err != nil {
		panic(err)
	}
}
