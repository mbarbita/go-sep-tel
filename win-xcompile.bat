set pre=tel-server-

set os=windows
set arch=386
set GOOS=%os%
set GOARCH=%arch%
echo compile: %os%-%arch%
go build -o %pre%%GOOS%-%GOARCH%.exe

set os=windows
set arch=amd64
set GOOS=%os%
set GOARCH=%arch%
echo compile: %os%-%arch%
go build -o %pre%%GOOS%-%GOARCH%.exe

set os=linux
set arch=386
set GOOS=%os%
set GOARCH=%arch%
echo compile: %os%-%arch%
go build -o %pre%%GOOS%-%GOARCH%

set os=linux
set arch=amd64
set GOOS=%os%
set GOARCH=%arch%
echo compile: %os%-%arch%
go build -o %pre%%GOOS%-%GOARCH%