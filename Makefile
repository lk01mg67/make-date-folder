.PHONY: build run clean test help

BINARY_NAME=make-date-folder

## build: 프로젝트 빌드
build:
	go build -o $(BINARY_NAME) main.go

## run: 프로그램 실행 (사용법: make run ARGS="./ 5")
run:
	go run main.go $(ARGS)

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
