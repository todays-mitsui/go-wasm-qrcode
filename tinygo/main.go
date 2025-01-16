package main

import (
	b64 "encoding/base64"
	"fmt"
	"unsafe"

	"github.com/skip2/go-qrcode"
)

func main() {
	c := make(chan struct{})
	<-c
}

var bufSize uint32

func stringToPtr(s string) unsafe.Pointer {
	buf := []byte(s)
	bufSize = uint32(len(buf))
	ptr := &buf[0]
	return unsafe.Pointer(ptr)
}

//go:wasmexport genQrcodePng
func genQrcodePng(data string) unsafe.Pointer {
	png, err := qrcode.Encode(data, qrcode.Medium, 800)
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	b64Png := b64.StdEncoding.EncodeToString(png)

	return stringToPtr(b64Png)
}

//go:wasmexport getBufSize
func getBufSize() uint32 {
	return bufSize
}
