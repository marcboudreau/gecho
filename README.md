# gecho
**gecho** is a binary that is useful for testing launch scripts.  All that it does
is output its environment variables and the command line arguments it was given
in a JSON representation.

By using **gecho**, instead of your regular binary, you can easily test that the
launch script sets up the correct environment variables and command line
arguments.This is especially useful when the launch script has multiple
conditional branches, or the target binary is cumbersome to launch.

## Usage
```console
./gecho
```
Will output `{"arguments":["./gecho"],"envvars":["NAME=VALUE"]}`.  The `envvars`
array will contain an element for every environment variable that was passed to
the process (i.e. exported on Unix-like platforms).  The elements are reported
as they are received from the `os.Environ()` function: variable name, equal
sign, and variable value as a single string.

```console
./gecho -la directory/
```
Will output `{"arguments":["./gecho","-la","directory/"],"envvars":[NAME=VALUE]}`.

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
