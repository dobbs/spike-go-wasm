# Spike Go and WASM

because I wanna play with running compiled things in the browser

# Usage

```
docker build -t spike-go-wasm .
docker run --name spike --rm -d -p 2015:2015 spike-go-wasm
open http://localhost:2015
# look at the console in your browser
# cleanup after you're done:
docker stop spike
```
