build target="echo":
    go build -ldflags "-s -w" http-{{target}}.go
