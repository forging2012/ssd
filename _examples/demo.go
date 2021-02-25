package main

/*
$ curl "http://localhost:9999"
Hello ssd
$ curl "http://localhost:9999/panic"
{"message":"Internal Server Error"}
$ curl "http://localhost:9999"
Hello ssd

>>> log
2020/01/09 01:00:10 Route  GET - /
2020/01/09 01:00:10 Route  GET - /panic
2020/01/09 01:00:22 [200] / in 25.364µs
2020/01/09 01:00:32 runtime error: index out of range
Traceback:
        /usr/local/Cellar/go/1.12.5/libexec/src/runtime/panic.go:523
        /usr/local/Cellar/go/1.12.5/libexec/src/runtime/panic.go:44
        /Users/7days-golang/day7-panic-recover/main.go:47
        /Users/7days-golang/day7-panic-recover/ssd/context.go:41
        /Users/7days-golang/day7-panic-recover/ssd/recovery.go:37
        /Users/7days-golang/day7-panic-recover/ssd/context.go:41
        /Users/7days-golang/day7-panic-recover/ssd/logger.go:15
        /Users/7days-golang/day7-panic-recover/ssd/context.go:41
        /Users/7days-golang/day7-panic-recover/ssd/router.go:99
        /Users/7days-golang/day7-panic-recover/ssd/ssd.go:130
        /usr/local/Cellar/go/1.12.5/libexec/src/net/http/server.go:2775
        /usr/local/Cellar/go/1.12.5/libexec/src/net/http/server.go:1879
        /usr/local/Cellar/go/1.12.5/libexec/src/runtime/asm_amd64.s:1338

2020/01/09 01:00:32 [500] /panic in 395.846µs
2020/01/09 01:00:38 [200] / in 6.985µs
*/

import (
	"log"
	"net/http"

	"github.com/forging2012/ssd"
)

func main() {
	r := ssd.Default()
	r.GET("/", func(c *ssd.Context) {
		c.String(http.StatusOK, "Hello ssd\n")
	})

	// index out of range for testing Recovery()
	r.GET("/panic", func(c *ssd.Context) {
		names := []string{"ssd"}
		c.String(http.StatusOK, names[100])
	})

	log.Fatal(r.Run("localhost:8080"))
}
