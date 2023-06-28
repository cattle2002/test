package protocol

import (
	"errors"
	"fmt"
	"strings"
)

type HelloHttp struct {
	HeaderLen uint32
	BodyLen   uint32
	Method    string
	uri       string
	Version   string
	Headers   map[string]string
	get       map[string]string
	post      map[string]string
}

func (this *HelloHttp) Length(httpMsg []byte) (msgLength uint64, err error) {
	fmt.Println("http的请求数据:", string(httpMsg))
	var httpMsgLength uint32 = uint32(len(httpMsg))
	if httpMsgLength == 0 {
		err = errors.New("接受http请求数据为空")
		return 0, err
	}
	//将字节数组转换成string
	httpMsgDataString := string(httpMsg)
	//查找\r\n结尾的字符串
	splitStrPos := strings.Index(httpMsgDataString, "\r\n\r\n")
	if splitStrPos == -1 || splitStrPos == 0 {
		//找不到的情况
		return 0, errors.New("http 请求报文不规范")
	}
	this.HeaderLen = uint32(splitStrPos) + 4
	contentLenStrPos := strings.Index(httpMsgDataString, "Content-Length: ")

	contentLenStr := httpMsgDataString[contentLenStrPos+len("Content-Length: "):]

	bodyLenStrPos := strings.IndexByte(contentLenStr, '\r')

	bodyLenStr := contentLenStr[0:bodyLenStrPos]
	fmt.Println("body长度为:", bodyLenStr)
	return uint64(len(bodyLenStr)), nil
}
