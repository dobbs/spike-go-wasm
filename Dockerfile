FROM golang:1.11.2 as src
WORKDIR /go/src/app
RUN printf "%s\n" > index.html \
  '<html><head><meta charset="utf-8">' \
  '<script src="wasm_exec.js"></script>' \
  '<script>' \
  'const go = new Go();' \
  'WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(' \
  '  result => {' \
  '    go.run(result.instance);' \
  '});' \
  '</script></head>' \
  '<body></body>' \
  '</html>'
  RUN cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
  COPY . .
  RUN go get -d -v .
  RUN GOOS=js GOARCH=wasm go build -o main.wasm

FROM abiosoft/caddy:0.11.1-no-stats
COPY --from=src \
  /go/src/app/main.wasm \
  /go/src/app/wasm_exec.js \
  /go/src/app/index.html \
  ./
