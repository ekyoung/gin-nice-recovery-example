package main

import (
    "github.com/ekyoung/gin-nice-recovery"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.New()      // gin.Default() installs gin.Recovery() so use gin.New() instead
    router.Use(gin.Logger()) // Install the default logger, not required

    // Install nice.Recovery, passing the name of the html template to render, and data to use
    router.Use(nice.Recovery("error.tmpl", gin.H{
        "title": "Error",
    }))

    // Load templates as usual
    router.LoadHTMLFiles("error.tmpl")

    router.GET("/", func(c *gin.Context) {
        panic("Doh!")
    })

    router.Run(":8080")
}
