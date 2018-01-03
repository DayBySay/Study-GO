run:
	go run Study-GO.go

fmt:
	gofmt -w *.go
	
save: fmt
	git add .
	git commit -m '...'
