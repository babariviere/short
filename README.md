# short

URL Shortener written in Go.

# Requirements

- [go 1.21.6 or more](https://go.dev/doc/install)
- [Docker](https://www.docker.com/)
- [atlas](https://atlasgo.io/getting-started)
- [sqlc](https://docs.sqlc.dev/en/stable/overview/install.html)
- [just](https://github.com/casey/just#installation) (optional)

> [!NOTE]
> If you are using [Nix](https://nixos.org/), you can just run `nix develop .` or `direnv allow` if you use Direnv.

## Dev dependencies

- [ogen](https://github.com/ogen-go/ogen#install) to generate OpenAPI server

# Running the app

Once you have installed all requirements, you can simply do:

```sh
docker compose up -d
just migrate # or atlas migrate apply --env local
go run .
```

## Accessing the docs

The docs is using the `openapi.yaml` provided in the root directory. A web UI is available at http://localhost:8080/docs.
It's using Rapidoc under the hood.

## Testing

You can simply test using:

```sh
# ensure postgres is up
docker compose up -d
go test ./...
```

It will create database on the fly for each test.
