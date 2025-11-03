# Mothalali

A tiny, toy Git-like project inspired by ugit.

## Requirements
- Python >= 3.6
- macOS/Linux shell (examples use zsh/bash)

## Quickstart

1) Create and activate a virtual environment
```bash
python3 -m venv .venv
source .venv/bin/activate
```

2) Install in editable mode
```bash
python -m pip install -e .
```

3) Run
- Installed console script (recommended):
```bash
mothalali
```
- Or as a module from the project root:
```bash
python -m mothalali.cli
```

## Project layout
```
.
├─ setup.py
├─ README.md
├─ mothalali/
│  ├─ __init__.py
│  └─ cli.py
```

## Development
- Reinstall after changing entry points or package metadata:
```bash
python -m pip install -e .
```
- Upgrade build tooling if needed:
```bash
python -m pip install --upgrade pip setuptools wheel
```

## Troubleshooting
- Command not found after install: ensure your venv is active and reinstall with `-e .`.
- Import errors from the repo root: use `python -m mothalali.cli` instead of `python mothalali/cli.py`.

## Reference
- Git Internals by Nikita Leshenko: [ugit](https://www.leshenko.net/p/ugit/#)

## License
MIT (or your preferred license).

