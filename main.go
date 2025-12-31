package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("사용법: go run main.go [경로] [생성일수]")
		fmt.Println("예시: go run main.go ./ 365")
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

	// 절대 경로로 변환 (선택 사항이나 권장됨)
	absPath, err := filepath.Abs(basePath)
	if err != nil {
		log.Fatalf("오류: 경로를 확인할 수 없습니다: %v", err)
	}

	now := time.Now()
	fmt.Printf("시작 경로: %s\n", absPath)
	fmt.Printf("%d일치 폴더 생성을 시작합니다...\n", count)

	successCount := 0
	for i := 1; i <= count; i++ {
		targetDate := now.AddDate(0, 0, i)
		year := targetDate.Format("2006")
		month := targetDate.Format("01")
		day := targetDate.Format("02")

		err := makeFolder(absPath, year, month, day)
		if err != nil {
			fmt.Printf("폴더 생성 실패 (%s-%s-%s): %v\n", year, month, day, err)
			continue
		}
		successCount++
	}

	fmt.Printf("완료! 총 %d개의 날짜 폴더가 생성되었습니다.\n", successCount)
}

func makeFolder(basePath, year, month, day string) error {
	// OS에 맞는 경로 구분자를 사용하도록 filepath.Join 사용
	path := filepath.Join(basePath, year, month, day)
	
	// 폴더 생성 (0755 권한)
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}
	return nil
}
