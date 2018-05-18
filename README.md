# snappy-cli
Simple CLI utility for the snappy compression format

Contains two complementary CLI utilities, one for compressing, the other
for uncompressing i.e `snappy-compress` and `snappy-uncompress`.

* Note:
Both utils:
- Accept directories (will operate on individual files within) and files
- Will remove original files

### Installation and usage

- snappy-compress

CLI utility to compress data in the snappy format.
It takes in either a source file or directory of source files e.g

```shell
$ go get github.com/Seventy-Two/snappy-cli/cmd/snappy-compress
$ snappy-compress filename
$ snappy-compress .
```

Will remove source file(s) and output files will have extension `.snappy`

- snappy-uncompress

CLI utility to uncompress data in the snappy format.
These files must have extension `.snappy` to avoid conflicts
It takes in either a source compressed file or directory of compressed files e.g

```shell
$ go get github.com/Seventy-Two/snappy-cli/cmd/snappy-uncompress
$ snappy-uncompress filename.snappy
$ snappy-uncompress .
```

Will remove source file(s) and output files will have identical names without `.snappy` extension
