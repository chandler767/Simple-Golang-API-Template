all:
	go build -o Simple_Golang_Api_Mac main.go
	GOOS=linux GOARCH=amd64 go build -o Simple_Golang_Api_Linux main.go
	GOOS=windows GOARCH=386 go build -o Simple_Golang_Api_Win32.exe main.go
