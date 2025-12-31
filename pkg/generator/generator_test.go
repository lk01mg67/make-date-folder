package generator

import (
	"os"
	"testing"
)

func TestGenerateFolders(t *testing.T) {
	// 임시 테스트용 디렉토리 생성
	tmpDir, err := os.MkdirTemp("", "generator_test")
	if err != nil {
		t.Fatalf("임시 디렉토리 생성 실패: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	count := 3
	result := GenerateFolders(tmpDir, count)

	if result.SuccessCount != count {
		t.Errorf("성공 횟수 불일치: 예상 %d, 실제 %d", count, result.SuccessCount)
	}

	if len(result.Errors) != 0 {
		t.Errorf("에러 발생: %v", result.Errors)
	}

	// 실제 폴더가 생성되었는지 확인 (첫 번째 날짜만 간단히 확인)
	// 참고: generator는 현재 시간 기준으로 내일부터 생성함
	// 여기서는 깊게 분기하지 않고 디렉토리 내 항목 개수만 확인
	entries, _ := os.ReadDir(tmpDir)
	// GenerateFolders는 년도별 폴더를 생성하므로 entries는 년도 폴더 하나(올해 또는 내년)여야 함
	if len(entries) == 0 {
		t.Error("폴더가 생성되지 않았습니다.")
	}
}
