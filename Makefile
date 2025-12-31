.PHONY: build run clean test help

BINARY_NAME=make-date-folder
SRC_PATH=./cmd/make-date-folder/main.go

## build: 프로젝트 빌드
build:
	go build -o $(BINARY_NAME) $(SRC_PATH)

## run: 프로그램 실행 (사용법: make run ARGS="./ 5")
run:
	go run $(SRC_PATH) $(ARGS)

## clean: 빌드된 바이너리 및 임시 폴더 삭제
clean:
	go clean
	rm -f $(BINARY_NAME)
	rm -rf test_folders

## test: 테스트 실행
test:
	go test ./...

## help: 도움말 출력
help:
	@echo "사용 가능한 명령:"
	@sed -n 's/^##//p' Makefile | column -t -s ':' |  sed -e 's/^/ /'
