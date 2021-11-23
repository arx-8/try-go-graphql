.PHONY: start
start:
	go run src/server.go

.PHONY: gqlgen-generate
gqlgen-generate:
	go run github.com/99designs/gqlgen generate
