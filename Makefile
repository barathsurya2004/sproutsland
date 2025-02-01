.PHONY : build

build: 
	@go build -o ./build/bin

.PHONY :run

run: build
	@./build/bin
