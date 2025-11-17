# Mothalali

A tiny Git-like toy inspired by Nikita Leshenko’s ugit, rewritten with Go + Cobra.

## Requirements
- Go 1.21+ (module targets Go 1.24)
- macOS/Linux shell (examples use zsh/bash)

## Quickstart

Run directly (no install):
```bash
go run . init
```

Create a reusable binary:
```bash
go build -o build/mothalali .
./build/mothalali init
```

Install globally (adds `mothalali` to your `$GOBIN` or `$GOPATH/bin`):
```bash
go install .
mothalali init
```

## Commands
- `mothalali init` – placeholder command demonstrating the CLI structure (currently prints “init called!!!”).
- Future commands (`add`, `commit`, …) are scaffolded in `cmd/root.go`.

## Project layout
```
.
├─ cmd/
│  ├─ init.go        # init sub-command
│  └─ root.go        # root command + scaffolding
├─ main.go           # entrypoint (invokes cmd.Execute)
├─ go.mod / go.sum
├─ README.md
└─ build/            # optional output from go build -o build/...
```

## Development
- Format / lint:
```bash
gofmt -w .
go vet ./...
```
- Test (none yet, but keep command handy):
```bash
go test ./...
```
- If you use cobra-cli helpers, ensure `go.mod` exists before running `cobra-cli init/add`.

## Reference
- Git Internals by Nikita Leshenko: [ugit](https://www.leshenko.net/p/ugit/#)

## License
MIT (or your preferred license).

