all:
	go build -o ./main.exe ./main.go

clean:
	del .\*.exe