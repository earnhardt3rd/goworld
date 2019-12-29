# goworld
Another Go Repo to say Hello

# Create a package
go build will trow away results, to use package use go install

go install will build and place package in %GOPATH%\pkg\windows_386\github.com\earnhardt3rd\goworld


-- add new package to main
c:\work\myGo\src\github.com\earnhardt3rd\goworld>type hello.go
package main

import (
   "fmt"
   "github.com/earnhardt3rd/goworld/string"
)

func main() {
   fmt.Println(string.Reverse("hello, world!"))
}

--build
c:\work\myGo\src\github.com\earnhardt3rd\goworld>go build

--run
c:\work\myGo\src\github.com\earnhardt3rd\goworld>goworld.exe
!dlrow ,olleh