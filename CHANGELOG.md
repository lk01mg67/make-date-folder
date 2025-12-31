# Changelog

All notable changes to this project will be documented in this file.

## [1.1.0] - 2025-12-31

### Added
- **프로젝트 표준화**: `Makefile`, `LICENSE` (MIT), `CHANGELOG.md` 추가
- **환경 설정**: macOS 및 Go 개발 환경에 최적화된 `.gitignore` 반영
- **GitHub 연동**: 저장소 Description 및 Topics 데이터 업데이트
- **문서화**: `README.md` 내 실제 실행 예시, 결과 메시지, 트리 구조 설명 추가

### Improved
- **OS 호환성**: `path/filepath` 패키지를 도입하여 Windows/macOS/Linux 통합 경로 처리 지원
- **안정성**: 
  - 입출력 인자 검증 및 `strconv` 에러 핸들링 추가
  - `filepath.Abs`를 통한 절대 경로 처리 로직 적용
  - 폴더 생성 실패 시 상세 에러 메시지 출력 및 작업 연속성 보장
- **UX**: 실행 시작 및 완료 시 요약 통계(성공 횟수 등) 메시지 출력 추가

## [1.0.0] - 2025-12-31

### Added
- 지정된 경로에 일수만큼 날짜별 폴더를 생성하는 기초 기능 구현 (Windows 전용)
