# x

Execute any script from anywhere on the path

## Install `x` with go

```sh
go install github.com/p-nerd/x@latest
```

## Commends

-   Root

```sh
x # it will run x.sh script, we can put argument that will be pass to the x.sh script (like this: x zip)
x -s <script name> # we can specify script with -s flag. it will run specified script. we also can put args (like this: x -s f.sh zip)
```

-   set

```sh
x set <script name> # change the default script name.
```

-   up

```sh
x up # run 'docker compose up' command on working path`
```

-   version

```sh
x version # shows tool version number`
```
