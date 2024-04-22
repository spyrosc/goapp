package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")
var numConn = flag.Int("connections", 100, "number of connections")

func main() {
	flag.Parse()
	log.SetFlags(0)

	wg := sync.WaitGroup{}

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/goapp/ws"}
	log.Printf("connecting to %s", u.String())

	done := make(chan struct{})
	exitChannel := make(chan os.Signal, 1)
	signal.Notify(exitChannel, syscall.SIGINT, syscall.SIGTERM)
	for i := 0; i < *numConn; i++ {

		c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			log.Fatal("dial:", err)
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			defer c.Close()
			for {
				select {
				default:
					_, message, err := c.ReadMessage()
					if err != nil {
						log.Println("read:", err)
						return
					}
					log.Printf("recv: %s", message)
				case <-done:
					return

				}
			}
		}()
	}

	<-exitChannel
	close(done)
	wg.Wait()
}
