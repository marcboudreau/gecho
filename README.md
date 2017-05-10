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
test gecho
echo $?
```
Returns `0` indicating a successful test.

## Building
```console
go build
```

## Testing
No tests yet.

## Installing
```console
go install github.com/marcboudreau/gecho
```
