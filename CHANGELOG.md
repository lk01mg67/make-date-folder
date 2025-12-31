# Changelog

All notable changes to this project will be documented in this file.

## [1.4.0] - 2025-12-31

### Added
- **단위 테스트**: `pkg/generator` 패키지에 대한 자동화된 단위 테스트 (`generator_test.go`) 추가

### Refactored
- **표준 Go 프로젝트 레이아웃 적용**:
  - `cmd/make-date-folder/main.go`: CLI 엔트리 포인트 (인자 처리 및 출력 담당)
  - `pkg/generator/generator.go`: 핵심 비즈니스 로직 (폴더 생성 라이브러리)
  - 로직 분리를 통해 재사용성 및 테스트 용이성 확보

## [1.3.0] - 2025-12-31

### Changed
- **프로젝트 구조 개편**: 메인 소스 코드를 `cmd/make-date-folder/` 폴더로 이동

## [1.2.0] - 2025-12-31

### Added
- **의존성 관리**: Go 버전을 1.17에서 1.23으로 업그레이드

## [1.1.0] - 2025-12-31

### Added
- **프로젝트 표준화**: `Makefile`, `LICENSE` (MIT), `CHANGELOG.md` 추가
- **환경 설정**: macOS 및 Go 개발 환경에 최적화된 `.gitignore` 반영
- **GitHub 연동**: 저장소 Description 및 Topics 데이터 업데이트
- **문서화**: `README.md` 내 실제 실행 예시 추가

### Improved
- **OS 호환성**: `path/filepath` 패키지 도입
- **안정성**: 에러 핸들링 및 절대 경로 처리 로직 적용
- **UX**: 실행 요약 통계 메시지 출력 추가

## [1.0.0] - 2025-12-31

### Added
- 지정된 경로에 일수만큼 날짜별 폴더를 생성하는 기초 기능 구현 (Windows 전용)
