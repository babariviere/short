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

# Running the app

Once you have installed all requirements, you can simply do:

```sh
docker compose up -d
just migrate # or atlas migrate apply --env local
go run .
```
