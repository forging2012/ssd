package main

import "ssd"

func main() {
	//r:=ssd.Default()
	r := ssd.New()
	r.Use(ssd.Recovery(), ssd.LoggerUUID())

	r.GET("/", func(c *ssd.Context) {
		c.JSON(200, "Hello world!")
	})
	r.Run(":8080")
}

// 中间件性能测试
// 打开：r.Use(ssd.Recovery(), ssd.Logger())
//➜  ssd git:(master) wrk10 http://localhost:8080/
//Running 10s test @ http://localhost:8080/
//  100 threads and 100 connections
//  Thread Stats   Avg      Stdev     Max   +/- Stdev
//    Latency     9.30ms   18.95ms 177.56ms   92.84%
//    Req/Sec   230.51
//   89.81     1.18k    72.67%
//  228355 requests in 10.06s, 26.79MB read
//Requests/sec:  22688.85
//Transfer/sec:      2.66MB

// 关闭：r.Use(ssd.Recovery(), ssd.Logger())
//➜  ssd git:(master) ✗ wrk10 http://localhost:8080/
//Running 10s test @ http://localhost:8080/
//  100 threads and 100 connections
//  Thread Stats   Avg      Stdev     Max   +/- Stdev
//    Latency     2.82ms    4.06ms 110.60ms   93.26%
//    Req/Sec   488.22    123.95     3.32k    87.32%
//  487446 requests in 10.10s, 57.18MB read
//Requests/sec:  48257.30
//Transfer/sec:      5.66MB

// ssd.LoggerUUID()
//➜  ssd git:(master) ✗ wrk10 http://localhost:8080/
//Running 10s test @ http://localhost:8080/
//  100 threads and 100 connections
//  Thread Stats   Avg      Stdev     Max   +/- Stdev
//    Latency    21.18ms   49.02ms 346.35ms   90.92%
//    Req/Sec   170.20     79.80   575.00     54.98%
//  160080 requests in 10.07s, 18.78MB read
//Requests/sec:  15900.19
//Transfer/sec:      1.87MB
