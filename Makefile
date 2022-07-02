all:
	go build -o ./main.exe ./main.go

heroku-bash:
	heroku run bash

heroku-push:
	git push heroku

push-amend:
	git add .
	echo "======== ADDED ========"
	git commit --amend --no-edit
	echo "======== AMENDED ========"
	git push origin master --force
	echo "======== PUSHED TO GITHUB ========"
	git push heroku --force
	echo "======== PUSHED TO HEROKU ========"

clean:
	del .\*.exe