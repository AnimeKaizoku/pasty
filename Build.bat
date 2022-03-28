@echo off
TITLE Build latest version
echo.
echo Building from branch:
git branch
echo.
echo Pulling latest version....
git pull
echo.
echo Building binary pasty.exe
go build -o pasty.exe ./cmd/pasty/main.go

