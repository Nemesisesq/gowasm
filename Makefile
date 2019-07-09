all: clean compile_wasm

clean:
	if [ -a lib.wasm ];  then rm lib.wasm; fi;

compile_wasm:
	@echo "compiling lib.wasm"
	GOARCH=wasm GOOS=js go build -o lib.wasm wasm.go