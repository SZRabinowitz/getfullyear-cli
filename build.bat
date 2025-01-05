@echo off
setlocal

:: Define the Go source file
set SOURCE_FILE=fullyear-cli.go

:: Windows
set GOOS=windows
set GOARCH=amd64
go build -o fullyear-cli-windows.exe %SOURCE_FILE%

:: Linux
set GOOS=linux
set GOARCH=amd64
go build -o fullyear-cli-linux %SOURCE_FILE%

:: macOS
set GOOS=darwin
set GOARCH=amd64
go build -o fullyear-cli-macos %SOURCE_FILE%

endlocal
echo Build complete.
