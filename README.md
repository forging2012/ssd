# README

## summary
`ssd` is a lightweight, idiomatic and composable framework for building Go HTTP services.

## install
`go get -u github.com/forging2012/ssd`

## example
See _examples/ for a variety of examples.

```go
package main

import "github.com/forging2012/ssd"

func main() {
	r := ssd.New()
	r.GET("/", func(c *ssd.Context) {
		c.JSON(200, "Hello world!")
	})
	r.Run(":8080")
}
```



# END
