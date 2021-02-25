package ssd

import (
	uuid "github.com/satori/go.uuid"
	"log"
	"time"
)

func LoggerUUID() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// uuid
		uid := uuid.NewV4()
		// Calculate resolution time
		log.Printf("[%d] [%s] %s in %v", c.StatusCode, uid, c.Request.RequestURI, time.Since(t))
	}
}
