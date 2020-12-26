build:
	go build

run:
	go run main.go

exec:
	./go-ocr

pull_master:
	git pull origin master

push_master:
	git push origin master

start:
	make build
	make exec