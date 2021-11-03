package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

type Connect struct {
	Buf  []byte
	Conn net.Conn
}

const (
	HeaderFlagLen = 3
	HeaderLen     = 7
	HeaderFlag1   = 0xaa
	HeaderFlag2   = 0xbb
	HeaderFlag3   = 0xcc
)

func (c *Connect) Header() (dataLen uint32) {
	if len(c.Buf) < HeaderLen { //不足协议头长度，下次处理
		return
	}
	if c.Buf[0] != HeaderFlag1 || c.Buf[1] != HeaderFlag2 || c.Buf[2] != HeaderFlag3 { //检测协议是否正确，如果有异常，丢弃错误数据
		var clear bool
		for i := 1; i < len(c.Buf)-2; i++ {
			if c.Buf[i] == HeaderFlag1 && c.Buf[i+1] == HeaderFlag2 && c.Buf[i+2] == HeaderFlag3 {
				c.Buf = c.Buf[i:]
				clear = true
			}
		}
		if clear == false {
			c.Buf = []byte{}
			return
		}
	}

	buff := bytes.NewBuffer(c.Buf[HeaderFlagLen:HeaderLen])
	binary.Read(buff, binary.BigEndian, &dataLen) //读取4个字节
	return
}

func (c *Connect) Handle() {
	for {
		l := c.Header()
		if l == 0 {
			break
		}
		if uint32(len(c.Buf)) < l+HeaderLen {
			break
		}
		//n, err := c.Conn.Write(c.Buf[HeaderLen : HeaderLen+l])
		//if err != nil {
		//	log.Println(n, err)
		//}
		log.Println(l, c.Buf[HeaderLen : HeaderLen+l])
		c.Buf = c.Buf[HeaderLen+l:]
	}
}

func handle(conn net.Conn) {
	Con := Connect{
		Conn: conn,
	}
	defer Con.Conn.Close() //关闭连接
	log.Println("Connect :", Con.Conn.RemoteAddr())
	data := make([]byte, 1024)

	for {
		//只要客户端没有断开连接，一直保持连接，读取数据
		n, err := Con.Conn.Read(data)
		//数据长度为0表示客户端连接已经断开
		if n == 0 {
			fmt.Printf("%s has disconnect", Con.Conn.RemoteAddr())
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		Con.Buf = append(Con.Buf, data[:n]...)
		Con.Handle()
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8899")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Start listen localhost:8899")
	for {
		//开始循环接收客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		//一旦收到客户端连接，开启一个新的gorutine去处理这个连接
		go handle(conn)
	}
}
