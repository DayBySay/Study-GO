run:
	go run Study-GO.go

save:
	gofmt -w *.go
	git add .
	git commit -m '...'
