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

func stringToPtr(s string) uint32 {
	buf := []byte(s)
	ptr := &buf[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	bufSize = uint32(len(buf))
	return uint32(unsafePtr)
}

//go:wasmexport genQrcodePng
func genQrcodePng(data string) uint32 {
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
