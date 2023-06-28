package main

import (
	"byteseq/protocol"
	"fmt"
	"net"
)

/* get请求携带请求体
C:\Users\gongzhaowei\AppData\Local\JetBrains\GoLand2023.1\tmp\GoLand\___go_build_serverhttp_go.exe
GET /user/login HTTP/1.1
User-Agent: PostmanRuntime-ApipostRuntime/1.1.0
Cache-Control: no-cache
content-type: application/json
Accept:
Accept-Encoding: gzip, deflate, br
Connection: keep-alive
Host: localhost:9501
Content-Length: 47
//get请求携带的请求体的参数
{
"no":"user",
"password":"123456"
}
*/

/* get请求不带请求体
GET /user/login HTTP/1.1
User-Agent: PostmanRuntime-ApipostRuntime/1.1.0
Cache-Control: no-cache
Accept:
Accept-Encoding: gzip, deflate, br
Connection: keep-alive
Host: localhost:9501

*/
/*  post请求携带请求体
POST /user/login HTTP/1.1
User-Agent: PostmanRuntime-ApipostRuntime/1.1.0
Cache-Control: no-cache
content-type: application/json
Accept: *
Accept-Encoding: gzip, deflate, br
Connection: keep-alive
Host: localhost:9501
Content-Length: 47

{
"no":"user",
"password":"123456"
}

*/

/* post请求不携带请求体
POST /user/login HTTP/1.1
User-Agent: PostmanRuntime-ApipostRuntime/1.1.0
Cache-Control: no-cache
Accept: *
Accept-Encoding: gzip, deflate, br
Connection: keep-alive
Host: localhost:9501
Content-Length: 0

*/
func main() {
	listen, err := net.Listen("tcp4", "localhost:9501")
	if err != nil {
		fmt.Println("listen error", err)
		return
	}
	conn, err := listen.Accept()
	if err != nil {
		fmt.Println("accept error", err)
		return
	}
	buf := make([]byte, 1024)
	http := protocol.HelloHttp{
		HeaderLen: 0,
		BodyLen:   0,
		Method:    "",
		Version:   "",
		Headers:   nil,
	}
	n, err := conn.Read(buf)
	//fmt.Println(n)
	length, err := http.Length(buf)
	fmt.Println(length)
	//if err != nil {
	//	fmt.Println("read error", err)
	//	return
	//}
	fmt.Println(string(buf[0:n]))
	//fmt.Print(buf[0:n])
	//fmt.Println("接收到的数据长度:", n)

}
