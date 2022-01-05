build target="echo":
    go build -o pg_flame -ldflags "-s -w" http-{{target}}.go
