package main

import (
	"log"
	"net"

	"mars-rover-photos/config"
	"mars-rover-photos/controllers"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	tmpl := iris.HTML("./views", ".html")
	tmpl.Reload(true)
	app.RegisterView(tmpl)

	app.HandleDir("/css", "./views/css")

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.ViewData("Message", "404 not found")
		ctx.View("header.html")
		ctx.View("error.html")
	})
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.ViewData("Message", "Internal Server Error")
		ctx.View("header.html")
		ctx.View("error.html")
	})

	// endpoints
	app.Get("/", controllers.GetPhotos)
	app.Get("/api/v1/mars-rover-photos", controllers.GetPhotos)

	// run listener
	listener, err := net.Listen("tcp", config.Server.Port)
	if err != nil {
		log.Printf("mars-rover-photos: server listening error: %v\n", err)
		return
	}
	app.Run(iris.Listener(listener))
}
