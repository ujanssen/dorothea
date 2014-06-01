test:
	go test -cover

cover:
	go test -coverprofile=coverage.out 
	go tool cover -html=coverage.out
	go tool cover -func=coverage.out

