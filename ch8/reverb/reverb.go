package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
	"time"
	"strings"
)

func main()  {
	const addr = "localhost:8000"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("listen tcp on:%s\n", addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}
		fmt.Println("get a connection")
		go handleConn(conn)
	}
}
func  handleConn(c net.Conn)  {
	input := bufio.NewScanner(c)
	for input.Scan(){
		// 如果不加go 的话 一个读取会阻塞住 下一个读取
		go echo(c, input.Text(), 1000 * time.Millisecond)
	}
	fmt.Println("close a connection")
	c.Close()
}

func  echo(c net.Conn, text string, delay time.Duration)  {
	text = strings.Trim(text, "\n`")
	if text != "" {
		fmt.Fprintf(c, "\t%s\n", strings.ToUpper(text))
		time.Sleep(delay)
		fmt.Fprintf(c, "\t%s\n", text)
		time.Sleep(delay)
		fmt.Fprintf(c, "\t%s\n", strings.ToLower(text))
	}

}

