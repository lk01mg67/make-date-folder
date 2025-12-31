package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Result contains information about the folder generation process
type Result struct {
	SuccessCount int
	Errors       []error
}

// GenerateFolders 지정된 경로에 오늘 이후의 날짜별 폴더를 생성합니다.
func GenerateFolders(basePath string, count int) Result {
	now := time.Now()
	result := Result{
		Errors: []error{},
	}

	for i := 1; i <= count; i++ {
		targetDate := now.AddDate(0, 0, i)
		year := targetDate.Format("2006")
		month := targetDate.Format("01")
		day := targetDate.Format("02")

		err := makeFolder(basePath, year, month, day)
		if err != nil {
			result.Errors = append(result.Errors, fmt.Errorf("%s-%s-%s 실패: %w", year, month, day, err))
			continue
		}
		result.SuccessCount++
	}

	return result
}

func makeFolder(basePath, year, month, day string) error {
	path := filepath.Join(basePath, year, month, day)
	return os.MkdirAll(path, 0755)
}
