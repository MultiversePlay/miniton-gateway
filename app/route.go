package app

import (
	"github.com/gin-gonic/gin"
	"www.miniton-gateway.com/app/controller/area"
	"www.miniton-gateway.com/app/controller/auth"
	"www.miniton-gateway.com/app/controller/demo"
	"www.miniton-gateway.com/app/controller/user"
	"www.miniton-gateway.com/pkg/http"
)

func Init() {
	demoRoute()
	tonAPIRoute()
}

func tonAPIRoute() {
	r := http.Server.Router
	apiGroup := r.Group("/api")
	userGroup := apiGroup.Group("/user")
	areaGroup := apiGroup.Group("/area")
	tonAPIUserRoute(userGroup)
	tonAPIAreaRoute(areaGroup)
}

func tonAPIAreaRoute(g *gin.RouterGroup) {
	g.GET("/country", area.CountryList)
}

func tonAPIUserRoute(g *gin.RouterGroup) {
	g.Use(auth.Auth())
	g.GET("/detail", user.Detail)
}

func demoRoute() {
	r := http.Server.Router
	d := r.Group("/demo")
	d.GET("/detail", demo.Detail)
	d.GET("/list", demo.List)
	d.POST("/create", demo.Create)
}
