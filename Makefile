.PHONY: all
all: test war gofish

.PHONY: test
test: cardlib/*.go
	go test -v cardlib/*.go

war: games/war/*.go
	go build games/war/war.go
	./war

gofish: games/gofish/*.go
	go build games/gofish/gofish.go
	./gofish

deps:
	go get .

clean:
	rm war gofish
