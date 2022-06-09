all:
	go build -o ./main.exe ./main.go

heroku-bash:
	heroku run bash

clean:
	del .\*.exe