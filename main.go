package main

import (
    "github.com/gin-gonic/gin"

    "github.com/ekyoung/gin-nice-recovery"
)

func main() {
    router := gin.New()      // gin.Default() installs gin.Recovery() so use gin.New() instead
    router.Use(gin.Logger()) // Install the default logger, not required

    // Install nice.Recovery, passing the handler to call after recovery
    router.Use(nice.Recovery(recoveryHandler))

    // Load templates as usual
    router.LoadHTMLFiles("error.tmpl")

    // Define routes as usual
    router.GET("/", func(c *gin.Context) {
        panic("Doh!")
    })

    router.Run(":8080")
}

func recoveryHandler(c *gin.Context, err interface{}) {
    c.HTML(500, "error.tmpl", gin.H{
        "title": "Error",
        "err":   err,
    })
}
