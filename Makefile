.PHONY: gocards
war: games/war/*.go
	go build games/war/war.go
