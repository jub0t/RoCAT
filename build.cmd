REM Create the build folder
mkdir ./build

REM Build Executable for Windows
go build -o ./build/rocat.exe

REM Set Compile Settings for Linux
set GOOS=linux
set GOARCH=amd64
set GOHOSTOS=linux

REM Build Executable for Linux
go build -o ./build/rocat
