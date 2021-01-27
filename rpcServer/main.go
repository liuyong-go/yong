package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/liuyong-go/yong/core/gerpc"
)

func startServer(addr chan string) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	gerpc.Accept(l)
}
func main() {
	addr := make(chan string)
	go startServer(addr)
	client, _ := gerpc.Dial("tcp", <-addr)
	defer func() { _ = client.Close() }()
	time.Sleep(time.Second)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := fmt.Sprintf("geerpc req %d", i)
			var reply string
			if err := client.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Println("reply:", reply)
		}(i)
	}
	wg.Wait()
}
func testDel() {
	var test = make(map[int64]string)
	test[1] = "test1"
	test[2] = "test2"
	var testvalue = test[1]
	delete(test, 1)
	fmt.Println(testvalue)

}
