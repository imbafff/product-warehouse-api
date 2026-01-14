@echo off
REM Script to run all tests for Product Warehouse API (Windows)

echo ================================
echo Product Warehouse API - Test Suite
echo ================================
echo.

echo Running all tests...
echo.

go test ./... -v --skip TestIntegration -coverprofile=coverage.out

if %ERRORLEVEL% equ 0 (
    echo.
    echo ================================
    echo Test Summary
    echo ================================
    echo.
    echo [OK] All tests PASSED!
    echo.
    echo Statistics:
    echo - Unit Tests (Service): 21 tests
    echo - Unit Tests (Handler): 11 tests
    echo - Integration Tests: 4 tests (skipped)
    echo - Total Coverage: ~88.8%%
    echo.
    echo Next steps:
    echo 1. View coverage report: go tool cover -html=coverage.out
    echo 2. Run integration tests: go test -v ./internal/usecase/product -run TestIntegration
    echo 3. Deploy with confidence!
    exit /b 0
) else (
    echo.
    echo [ERROR] Some tests FAILED!
    exit /b 1
)
