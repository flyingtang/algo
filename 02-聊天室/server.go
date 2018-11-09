package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type Client struct {
	C    chan string
	Name string
	Addr string
}

var onLineMap sync.Map
var message = make(chan string)

func broadcast() {
	for {
		msg := <-message
		onLineMap.Range(func(_, value interface{}) bool {
			c := value.(Client)
			c.C <- msg
			return true
		})
	}
}

func main() {
	l, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err.Error())
	}
	defer l.Close()
	// 广播
	go broadcast()
	//  主循环
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}

func makeMessage(cli Client, msg string) string {
	return "[" + cli.Addr + "]" + cli.Name + ": " + msg + "\n"
}

func handleConn(conn net.Conn) {
	clientAddr := conn.RemoteAddr().String()
	cli := Client{make(chan string), clientAddr, clientAddr}
	onLineMap.Store(clientAddr, cli)
	go writeMsg(cli, conn)
	message <- makeMessage(cli, "login")
	isQuit := make(chan bool)
	hasData := make(chan bool)
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if n == 0 { // 对方出问题
				isQuit <- true
				fmt.Println("err " + err.Error())
				return
			}
			message <- makeMessage(cli, string(buf[:n]))
		}
		hasData <- true
	}()
	for {
		select {
		case <-isQuit:
			onLineMap.Delete(clientAddr)
			message <- makeMessage(cli, "login out")
			return
		case <-hasData:
		case <-time.After(time.Second * 10):
			onLineMap.Delete(clientAddr)
			message <- makeMessage(cli, "login out")
			return
		}
	}
}

func writeMsg(cli Client, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg))
	}
}
