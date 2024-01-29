alias b := build
alias r := run
alias l := lint

build:
  go build .

run:
  go run .

lint:
  golangci-lint run

gen:
  go generate ./...

sqlc:
  sqlc generate
  sqlc vet

gen-migration NAME:
  atlas migrate diff {{NAME}} --env local
  atlas migrate apply --env local

migrate:
  atlas migrate apply --env local
