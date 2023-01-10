package main

import (
	"html/template"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	r.Static("/assets", "frontend/assets")
	r.LoadHTMLGlob("frontend/templates/*.html")

	// Web pages
	r.GET("/", IndexPage)

	// API endpoints
	r.GET("/api/ping", PingEndpoint)

	r.Run("localhost:8080")
}
