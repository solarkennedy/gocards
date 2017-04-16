.PHONY: test
test:
	go test -v cardlib/*.go

war: games/war/*.go
	go build games/war/war.go

gofish: games/gofish/*.go
	go build games/gofish/gofish.go

