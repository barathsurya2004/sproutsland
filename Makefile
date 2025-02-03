.PHONY : build

build: 
	@mkdir -p build
	@go build -o ./build/bin

.PHONY :run

run: build
	@./build/bin

.PHONY : audit 

audit:
	@go mod tidy
	@go mod verify


