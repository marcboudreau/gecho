# gecho
**gecho** is a binary that is useful for testing launch scripts.  All that it does
is output its environment variables and the command line arguments it was given
in a JSON representation.

By using **gecho**, instead of another binary, a test can be constructed that
runs the launch script and validates that the correct arguments and environment
variables were set.  This is especially useful when the launch script has multiple
conditional branches, or the target binary is cumbersome to launch.

## Usage
```console
./gecho
```
Will output `{"arguments":["./gecho"]}`.

```console
./gecho -la directory/
```
Will output `{"arguments":["./gecho","-la","directory/"]}`.

## Building
```console
go build
```

## Testing
```console
go test -v
```

## Installing
```console
go install github.com/marcboudreau/gecho
```
