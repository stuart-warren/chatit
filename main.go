package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"text/template"
)

var addr = flag.String("addr", ":5000", "main http address")
var rootTempl = template.Must(template.ParseFiles("index.html"))

func rootHandler(context *gin.Context) {
	r := context.Request
	w := context.Writer
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	rootTempl.Execute(w, r.Host)
}

func main() {
	flag.Parse()
	log.Println("Starting.")
	go h.run()
	r := gin.Default()
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.File("favicon.ico")
	})
	r.GET("/", rootHandler)
	r.GET("/ws", wsHandler)

	err := http.ListenAndServe(*addr, r)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
