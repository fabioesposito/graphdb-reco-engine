package main

import (
	"fmt"
	"html/template"
	"strings"

	h "graphdb-reco-engine/handlers"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("init")
}

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	r.Static("/assets", "html/assets")
	r.LoadHTMLGlob("html/templates/*.html")

	// App Routes
	// Web pages
	r.GET("/", h.IndexPage)

	// API endpoints
	r.GET("/api/ping", h.PingEndpoint)

	r.Run("localhost:8080")
}
