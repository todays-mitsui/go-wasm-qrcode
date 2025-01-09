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

//go:wasmexport getBuffer
func getBuffer(length uint32) *byte {
	buf := make([]byte, length)
	return &buf[0]
}

var bufSize uint32

func stringToPtr(s string) (uint32, uint32) {
	buf := []byte(s)
	ptr := &buf[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return uint32(unsafePtr), uint32(len(buf))
}

//go:wasmexport genQrcodePng
func genQrcodePng(data string) uint32 {
	png, err := qrcode.Encode(data, qrcode.Medium, 800)
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	b64Png := b64.StdEncoding.EncodeToString(png)

	ptr, size := stringToPtr(b64Png)
	bufSize = size
	return ptr
}

//go:wasmexport getBufSize
func getBufSize() uint32 {
	return bufSize
}
