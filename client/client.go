package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	// 连接服务端
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	// 发送数据
	dataToSend := uint32(5678)
	buffer := make([]byte, 4)
	binary.LittleEndian.PutUint32(buffer, dataToSend)

	_, err = conn.Write(buffer)
	if err != nil {
		fmt.Println("Error sending data:", err.Error())
		return
	}

	fmt.Println("Data sent.")

	// 接收响应数据
	responseBuffer := make([]byte, 4)
	_, err = conn.Read(responseBuffer)
	if err != nil {
		fmt.Println("Error receiving response:", err.Error())
		return
	}

	// 将接收到的字节序列转换为大端整数
	receivedResponse := binary.BigEndian.Uint32(responseBuffer)
	fmt.Printf("Received Response: %d\n", receivedResponse)
}
