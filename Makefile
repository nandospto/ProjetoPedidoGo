run-api:
	cls
	go build -o bin/api.exe ./app/main.go
	./bin/api.exe