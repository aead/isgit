## isgit(1)

`isgit` is a small command-line tool that consumes a list of file paths and outputs all paths
that point to a git repository.

```
Usage:
    isgit [-o OUTPUT] [PATH ...]
 
Options:
    -o, --output OUTPUT         Write the result to the file at path OUTPUT.

isgit iterates over all PATH arguments and writes any PATH that
points to a git repository (containing a '.git' subdirectory) to
OUTPUT.
isgit reads a list of file paths from standard input if no PATH
argument(s) have been provided or when one PATH is '-'.

OUTPUT defaults to standard output.
```

### Example

Filter all paths that are git repositories:
```sh
isgit $HOME/project $HOME/my-repo
```

Find all git repositories under `$HOME` using [`fd`](https://github.com/sharkdp/fd) or GNU `find`:
```sh
fd -HL -t d "."  "$HOME/go" | isgit
```
```sh
find -L "$HOME" -type d | isgit
```
> The output of both commands may differ because `fd` honors ignore files. See `fd --help`

Fuzzy-search all git repositories under `$HOME` using [`fzf`](https://github.com/junegunn/fzf)
and [`fd`](https://github.com/sharkdp/fd) and a show commit history preview:
```sh
fd -H -L -t d "."  "$HOME" | isgit | fzf \
  --height 100% --reverse --border --preview-window right:50% \
  --preview 'git -C {} log --color=always --pretty=oneline --abbrev-commit'
``` 

### Install

#### Binary Releases

| OS      | Arch  | Binary                                                                                                    |
|---------|-------|-----------------------------------------------------------------------------------------------------------|
| Linux   | amd64 | [isgit-linux-amd64](https://github.com/aead/isgit/releases/latest/download/isgit-linux-amd64)             |
| MacOS   | amd64 | [isgit-darwin-amd64](https://github.com/aead/isgit/releases/latest/download/isgit-darwin-amd64)           |
| Windows | amd64 | [isgit-windows-amd64.exe](https://github.com/aead/isgit/releases/latest/download/isgit-windows-amd64.exe) |
| FreeBSD | amd64 | [isgit-freebsd-amd64](https://github.com/aead/isgit/releases/latest/download/isgit-freebsd-amd64)         |
|         |       |                                                                                                           |
| Linux   | arm   | [isgit-linux-arm](https://github.com/aead/isgit/releases/latest/download/isgit-linux-arm)                 |
| Linux   | arm64 | [isgit-linux-arm64](https://github.com/aead/isgit/releases/latest/download/isgit-linux-arm64)             |

#### Download via cURL

| OS      | Arch  | cURL Command                      |
|---------|-------|-----------------------------|
| Linux   | amd64 | `curl -SL --tlsv1.2 -o isgit https://github.com/aead/isgit/releases/latest/download/isgit-linux-amd64`           |
| MacOS   | amd64 | `curl -SL --tlsv1.2 -o isgit https://github.com/aead/isgit/releases/latest/download/isgit-darwin-amd64`          |
| Windows | amd64 | `curl -SL --tlsv1.2 -o isgit.exe https://github.com/aead/isgit/releases/latest/download/isgit-windows-amd64.exe` |
| FreeBSD | amd64 | `curl -SL --tlsv1.2 -o isgit https://github.com/aead/isgit/releases/latest/download/isgit-freebsd-amd64`         |
|         |       |                                                                                                                  |
| Linux   | arm   | `curl -SL --tlsv1.2 -o isgit https://github.com/aead/isgit/releases/latest/download/isgit-linux-arm`             |
| Linux   | arm64 | `curl -SL --tlsv1.2 -o isgit https://github.com/aead/isgit/releases/latest/download/isgit-linux-arm64`           |

#### From Source

```sh
go get github.com/aead/isgit
```

### License

Use of `isgit` is governed by the MIT license that can be found in the LICENSE file.
