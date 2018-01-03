run:
	go run Study-GO.go

fmt:
	gofmt -w *.go
	
add-commit: fmt
	git add .
	git commit -m '...'
