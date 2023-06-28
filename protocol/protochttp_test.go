package protocol

import (
	"fmt"
	"strings"
	"testing"
)

func TestSubstring(t *testing.T) {
	str := "asdfghjkl\r\n\r\n"
	index := strings.Index(str, "\r\n")
	fmt.Println(index)
}
