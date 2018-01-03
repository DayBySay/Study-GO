run:
	go run Study-GO.go

test:
	go test -v $$(go list ./...)

fmt:
	gofmt -w ./
	
add-commit: fmt
	git add .
	git commit -m '...'
