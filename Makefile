build:
	env GOOS=linux CGO_ENABLED=0 go build -o bin/getconfig lambda/getconfig/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy

.PHONY: build deploy clean