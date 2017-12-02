.PHONY: all
all: war gofish crazyeights

.PHONY: test
test: cardlib/*.go
	go test -v cardlib/*.go

war: games/war/*.go
	go build games/war/war.go
	./war

gofish: games/gofish/*.go
	go build games/gofish/gofish.go
	./gofish

crazyeights: games/crazyeights/*.go
	go build games/crazyeights/crazyeights.go
	./crazyeights

deps:
	dep ensure

clean:
	rm war gofish crazyeights
