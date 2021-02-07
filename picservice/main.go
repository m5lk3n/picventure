package main

import (
	"fmt"
	_ "image/png"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nathany/bobblehat/sense/screen"
	"github.com/nathany/bobblehat/sense/screen/texture"

	log "github.com/sirupsen/logrus"
)

// ScreenClearHandler clears the display of the Pi's Sense HAT
func ScreenClearHandler(c *gin.Context) {
	screen.Clear()

	c.JSON(http.StatusOK, gin.H{"message": "cleared", "status": http.StatusOK})
}

// ScreenDrawHandler draws the given pic on the Pi's Sense HAT
func ScreenDrawHandler(c *gin.Context) {
	pic := c.Param("pic")
	if pic == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "malformed request", "status": http.StatusBadRequest})
	}

	wd, _ := os.Getwd()
	pic = fmt.Sprintf("%s/pics/%s.png", wd, pic) // pictures come from pics subfolder, PNGs only

	tx, err := texture.Load(pic)
	if err != nil {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "requested resource not found", "status": http.StatusNotFound})
	}

	// draw PNG, first 8x8, starting at 0,0
	fb := screen.NewFrameBuffer()
	texture.Blit(fb.Texture, 0, 0, tx, 0, 0, 8, 8)
	screen.Draw(fb)

	c.JSON(http.StatusOK, gin.H{"message": "drawn", "status": http.StatusOK})
}

// NotFoundHandler is to indicate that requested resource could not be found
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

	v1 := router.Group("/api/v1/screen")
	{
		v1.GET("/clear", ScreenClearHandler)
		v1.GET("/draw/:pic", ScreenDrawHandler)
	}

	return router
}

func main() {
	router := SetupRouter()

	log.Infoln("picservice start...")
	defer log.Infoln("picservice shutdown!")

	// set port via PORT environment variable
	router.Run() // default port is 8080
}
