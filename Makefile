all:
	go build -o ./main.exe ./main.go

heroku-bash:
	heroku run bash

heroku-push:
	git push heroku

clean:
	del .\*.exe