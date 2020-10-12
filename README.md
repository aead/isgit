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

### License

Use of `isgit` is governed by the MIT license that can be found in the LICENSE file.

