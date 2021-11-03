package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

func main(){
	conn, err := net.Dial("tcp", ":8899")  //连接服务端
	if err != nil{
		fmt.Println(err)
		return
	}
	defer conn.Close()
	fmt.Println("Connect to localhost:8899 success")
	for i := 0; i < 100; i++{
		data := fmt.Sprintf(`{"index":%d, "name":"maqian", "age":21, "company":"intely"}`, i + 1)
		buff := bytes.Buffer{}
		binary.Write(&buff, binary.BigEndian, []byte{0xaa, 0xbb, 0xcc})
		binary.Write(&buff, binary.BigEndian, uint32(len(data)))
		log.Println(len(data), buff.Bytes())
		buff.WriteString(data)
		n, err := conn.Write(buff.Bytes())
		if err != nil{
			fmt.Println(err)
			continue
		}
		fmt.Printf("Send %d byte data : %s", n, data)
	}

	for{
		//一直循环读入用户数据，发送到服务端处理
		fmt.Print("Please input send data :")
		var a string
		fmt.Scan(&a)
		if a == "exit"{break}  //添加一个退出机制，用户输入exit，退出
		_, err := conn.Write([]byte(a))
		if err != nil{
			fmt.Println(err)
			return
		}
		data := make([]byte, 2048)
		n, err := conn.Read(data)
		if err != nil{
			fmt.Println(err)
			continue
		}
		fmt.Println("Response data :", string(data[:n]))
	}
}