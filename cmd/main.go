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
	r.Static("/assets", "public_html/assets")
	r.LoadHTMLGlob("public_html/templates/*.html")

	// App Routes
	// Web pages
	r.GET("/", h.IndexPage)

	// API endpoints
	r.GET("/api/ping", h.PingEndpoint)

	r.Run("localhost:8080")
}
