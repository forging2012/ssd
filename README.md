# README

## summary
ssd is a web framework written in Go (Golang), inspired from gin framework.

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

## Benchmarks

```text
➜  ~ wrk -t 100 -c 100 -d 10s http://localhost:8080/
Running 10s test @ http://localhost:8080/
  100 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.32ms    7.25ms 144.32ms   90.68%
    Req/Sec   427.33    128.05     1.24k    72.61%
  426625 requests in 10.04s, 50.04MB read
Requests/sec:  42509.09
Transfer/sec:      4.99MB
```


# END
