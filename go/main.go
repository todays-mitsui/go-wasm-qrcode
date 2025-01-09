package main

import (
	b64 "encoding/base64"
	"fmt"
	"syscall/js"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("genQrcodePng", js.FuncOf(genQrcodePng))

	<-c
}

func genQrcodePng(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		fmt.Println("No argument received")
		return nil
	}
	data := args[0].String()

	png, err := qrcode.Encode(data, qrcode.Medium, 800)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return b64.StdEncoding.EncodeToString(png)
}
