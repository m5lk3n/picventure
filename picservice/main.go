package main

import (
	_ "image/png"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nathany/bobblehat/sense/screen"
	"github.com/nathany/bobblehat/sense/screen/texture"

	log "github.com/sirupsen/logrus"
)

// ScreenClearHandler TODO
func ScreenClearHandler(c *gin.Context) {
	screen.Clear()

	c.JSON(http.StatusOK, gin.H{"message": "cleared", "status": http.StatusOK})
}

// ScreenDrawHandler TODO
func ScreenDrawHandler(c *gin.Context) {

	wd, _ := os.Getwd()
	pic := wd + "/pics/key.png"

	tx, err := texture.Load(pic)
	if err != nil {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "requested resource not found", "status": http.StatusNotFound})
	}
	fb := screen.NewFrameBuffer()
	texture.Blit(fb.Texture, 0, 0, tx, 0, 0, 8, 8) // 8x8, starting at 0,0
	screen.Draw(fb)

	c.JSON(http.StatusOK, gin.H{"message": "drawn", "status": http.StatusOK})
}

// LivenessHandler TODO
func LivenessHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "alive", "status": http.StatusOK})
}

// ReadinessHandler TODO
func ReadinessHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ready", "status": http.StatusOK})
}

// NotFoundHandler to indicate that requested resource could not be found
func NotFoundHandler(c *gin.Context) {
	// log this event as it could be an attempt to break in...
	log.Infoln("Not found, requested URL path:", c.Request.URL.Path)
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "requested resource not found", "status": http.StatusNotFound})
}

// SetupRouter is published here to allow setup of tests
func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.Use(gin.Recovery()) // "recover from any panics", write 500 if any

	router.NoRoute(NotFoundHandler)

	// public, generic API
	router.GET("/healthy", LivenessHandler)
	router.GET("/ready", ReadinessHandler)

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/screen/clear", ScreenClearHandler)
		v1.GET("/screen/draw", ScreenDrawHandler)
	}

	return router
}

func main() {
	router := SetupRouter()

	log.Infoln("web app start...")
	defer log.Infoln("web app shutdown!")

	// set port via PORT environment variable
	router.Run() // default port is 8080
}
