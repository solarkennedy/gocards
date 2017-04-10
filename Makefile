.PHONY: gocards
gocards: *.go
	go build .
	./gocards
