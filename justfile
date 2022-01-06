build target="echo":
    go build -o pg_flame -ldflags "-s -w" http-{{target}}.go
    tar zcvf pg_flame.tar.gz pg_flame assets
    rm -f pg_flame
