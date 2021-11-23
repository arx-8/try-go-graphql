.PHONY: start
start:
	go run server.go

.PHONY: gqlgen-generate
gqlgen-generate:
	go run github.com/99designs/gqlgen generate
