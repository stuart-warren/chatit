package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "time"
)

var (
    BLANK_TIME = time.Time{}
)

func messageGet(c *gin.Context) {
    ts := time.Now().UTC()
    msg := Message{Room: "#test", Timestamp: ts, Message: "This is test"}
    c.JSON(200, msg)
}

func messagePost(c *gin.Context) {
    var msg Message
    if c.Bind(&msg) {
        if msg.Timestamp == BLANK_TIME {
            msg.Timestamp = time.Now().UTC()
        }
        if msg.ClientId == "" {
            msg.ClientId = GetMD5Hash(c.Request.UserAgent())
        }
        if msg.Id == "" {
            msg.Id = GetUUID()
        }
        c.JSON(200, msg)
    } else {
        c.JSON(400, gin.H{"error":"400 Bad request"})
    }
}

func main() {
    log.Println("Starting.")
    r := gin.Default()
    r.GET("/message", messageGet)
    r.POST("/message", messagePost)

    http.Handle("/", r)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }
}
