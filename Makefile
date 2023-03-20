all:
	go build -o ./main.exe ./main.go

push-amend:
	git add .
	echo "======== ADDED ========"
	git commit --amend --no-edit
	echo "======== AMENDED ========"
	git push origin master --force
	echo "======== PUSHED TO GITHUB ========"

clean:
	del .\*.exe