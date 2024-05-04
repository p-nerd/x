# x

This is a command line to automate my software development workflow

## Install `x` with go

```sh
go install github.com/p-nerd/x@latest
```

## Commands

```
<...args>                   run the default script (default script: x.sh).
-s <script name> <...args>  run the specified script.
set <script name>           change the default script.
up                          run the 'docker compose up' command in the current working path.
commit                      run the 'git add -A; git commit -m <args>' in on command
push | git-all              run the 'git add -A; git commit -m <args>; git push' in on command
tree                        run the 'tree --gitignore' in on command
help | --help | -h          display this help information (also supports '--help').
version                     show the version number of the tool.
```
