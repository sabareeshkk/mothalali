# Mothalali

A tiny, toy Git-like project inspired by ugit, now implemented in Go.

## Requirements
- Go 1.21+
- macOS/Linux shell (examples use zsh/bash)

## Quickstart

### Run without installing
```bash
go run ./cmd/mothalali init
```

### Build a binary
```bash
go build -o build/mothalali ./cmd/mothalali
./build/mothalali init
```

The command initializes a `.mothalali` directory in the current working directory (if it does not already exist) and prints its absolute path, mimicking the original Python behavior.

## Project layout
```
.
├─ cmd/
│  └─ mothalali/
│     └─ main.go          # CLI entrypoint
├─ internal/
│  └─ mothalali/
│     └─ repo.go          # Repository helpers
├─ go.mod
├─ README.md
└─ build/                 # Created after go build
```

## Development
- Format code:
```bash
gofmt -w cmd/mothalali internal/mothalali
```
- Run go vet / tests (none yet):
```bash
go vet ./...
go test ./...
```

## Reference
- Git Internals by Nikita Leshenko: [ugit](https://www.leshenko.net/p/ugit/#)

## License
MIT (or your preferred license).

