package main

import (
	"fmt"
	"log"
	"make-date-folder/pkg/generator"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const version = "1.3.0"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	arg1 := os.Args[1]
	if arg1 == "-v" || arg1 == "--version" || arg1 == "version" {
		fmt.Printf("make-date-folder version %s\n", version)
		return
	}

	if len(os.Args) < 3 {
		printUsage()
		return
	}

	basePath := os.Args[1]
	countStr := os.Args[2]

	count, err := strconv.Atoi(countStr)
	if err != nil {
		log.Fatalf("오류: 생성 일수는 숫자여야 합니다: %v", err)
	}

	if strings.TrimSpace(basePath) == "" {
		log.Fatal("오류: 유효한 경로를 입력해 주세요.")
	}

	absPath, err := filepath.Abs(basePath)
	if err != nil {
		log.Fatalf("오류: 경로를 확인할 수 없습니다: %v", err)
	}

	fmt.Printf("시작 경로: %s\n", absPath)
	fmt.Printf("%d일치 폴더 생성을 시작합니다...\n", count)

	result := generator.GenerateFolders(absPath, count)

	for _, err := range result.Errors {
		fmt.Printf("오류 발생: %v\n", err)
	}

	fmt.Printf("완료! 총 %d개의 날짜 폴더가 생성되었습니다.\n", result.SuccessCount)
}

func printUsage() {
	fmt.Println("사용법: go run main.go [경로] [생성일수]")
	fmt.Println("옵션:")
	fmt.Println("  -v, --version  버전 정보 출력")
	fmt.Println("예시: go run main.go ./ 365")
}
