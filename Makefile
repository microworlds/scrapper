compile:
	GOOS=windows GOARCH=386 go build -o bin/scrapper-windows-386 scrapper.go
	GOOS=linux GOARCH=386 go build -o bin/scrapper-linux-386 scrapper.go
	GOOS=darwin GOARCH=386 go build -o bin/scrapper-darwin-386 scrapper.go