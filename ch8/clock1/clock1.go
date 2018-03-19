package main

import (
	"net"
	"log"
	"io"
	"time"
	"fmt"
)

func main()  {
	const addr = "localhost:8000"
	// 监听
	listener, err := net.Listen("tcp", addr)
	// 接受处理
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("listen tcp on:%s\n", addr)
	for {
		conn, err := listener.Accept()
		if err != nil{
			log.Print(err)
			continue
		}
		fmt.Println("get a connection")
		// 如果不加入go来将连接放入一个gorouting 则一个新的连接到来的时候会变得阻塞
		go handleConn(conn)
	}
}
func handleConn(conn net.Conn)  {
	// 定时向客户端写入当前的时间
	defer conn.Close()
	for  {
		_, err := io.WriteString(conn, time.Now().Format("13:04:05\n"))
		if err != nil {
			io.WriteString(conn, "server 500\n")
		}
		time.Sleep(1 * time.Second)
	}

}
