# lsgo

A lightweight terminal file manager written in Go.

Fast, minimal, and vim-inspired navigation for your filesystem.

---

## Features

- Vim-like navigation (`j`, `k`)
- Enter directories (`l`)
- Go to parent directory (`h`)
- Open files in default editor
- Scrollable viewport with offset handling
- Terminal-aware rendering (auto height detection)

---

## Installation

### Option 1: go install (recommended)

```bash
go install .
```

Make sure `$HOME/go/bin` is in your PATH:

```bash
export PATH="$HOME/go/bin:$PATH"
```

---

### Option 2: manual build

```bash
go build -o lsgo
sudo mv lsgo /usr/local/bin/
```

---

## Usage

Run in current directory:

```bash
lsgo
```

Or specify path:

```bash
lsgo --path=/your/directory
```

---

## Keybindings

| Key | Action                      |
| --- | --------------------------- |
| j   | move down                   |
| k   | move up                     |
| l   | enter directory / open file |
| h   | go to parent directory      |
| q   | quit                        |

---

## Configuration

On first run, `lsgo` creates a config file in:

```
~/.config/lsgo/config.json
```

You can override config path manually if needed.

Example config:

```json
{
  "standart_editor": "nvim"
}
```

---

## Requirements

- Go 1.20+
- Unix-like OS (Linux / macOS recommended)

---

## Roadmap

-

---

## Notes

This project is a learning-focused implementation of a terminal file manager. Inspired by tools like `nnn`, `ranger`, and `yazi`.

---

## License

MIT

