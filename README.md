# Google Chat TUI

Early-stage Go module containing authentication and configuration packages for a
Google Chat terminal application.

## Current scope

The tracked implementation is under [`internal/`](internal/). There is no published
`main` package or executable entry point at the repository root, so `go run .` is
not a supported launch command yet.

## Development

Use the Go version declared in [`go.mod`](go.mod), with dependencies pinned by
[`go.sum`](go.sum).

```sh
go list ./...
go test ./...
```

Add executable setup and authentication instructions with the command entry point.
Do not commit tokens or account credentials.
