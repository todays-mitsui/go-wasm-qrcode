all: go.wasm tinygo.wasm

go.wasm: go/main.go
	GOOS=js GOARCH=wasm go build -o app/wasm/go.wasm go/main.go

tinygo.wasm: tinygo/main.go
	GOOS=js GOARCH=wasm tinygo build -o app/wasm/tinygo.wasm tinygo/main.go

clean:
	rm -f app/wasm/go.wasm app/wasm/tinygo.wasm

dev:
	open http://localhost:8888/go.html \
	&& cd app \
	&& deno run --allow-env --allow-read --allow-net https://js.sabae.cc/liveserver.js 8888
