package main

import (
	"github.com/forging2012/ssd"
)

func main() {
	//  engine := New()
	//	engine.Use(Logger(), Recovery())
	r := ssd.Default()
	r.GET("/", func(c *ssd.Context) {
		c.JSON(200, "Hello world!")
	})
	r.Run(":8080")
}
