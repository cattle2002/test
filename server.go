package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	// 启动服务端监听
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Server started. Waiting for connections...")

	// 接受客户端连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err.Error())
		return
	}
	defer conn.Close()

	fmt.Println("Client connected.")

	// 接收数据
	buffer := make([]byte, 4) // 32位整数为4个字节
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading data:", err.Error())
		return
	}

	// 将接收到的字节序列转换为大端整数
	receivedData := binary.LittleEndian.Uint32(buffer)
	fmt.Printf("Received Data: %d\n", receivedData)

	// 发送响应数据
	responseData := uint32(1234)
	responseBuffer := make([]byte, 4)
	binary.BigEndian.PutUint32(responseBuffer, responseData)

	_, err = conn.Write(responseBuffer)
	if err != nil {
		fmt.Println("Error sending response:", err.Error())
		return
	}

	fmt.Println("Response sent.")
}
