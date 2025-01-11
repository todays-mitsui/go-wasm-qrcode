all: go.wasm tinygo.wasm

go.wasm: go/main.go
	GOOS=js GOARCH=wasm go build -o public/wasm/go.wasm go/main.go

tinygo.wasm: tinygo/main.go
	GOOS=js GOARCH=wasm tinygo build -o public/wasm/tinygo.wasm tinygo/main.go

clean:
	rm -f public/wasm/go.wasm public/wasm/tinygo.wasm

dev:
	open http://localhost:8888/go.html \
	&& cd public \
	&& deno run --allow-env --allow-read --allow-net https://js.sabae.cc/liveserver.js 8888
