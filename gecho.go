package main

import (
  "bytes"
  "fmt"
  "encoding/json"
  "os"
)

// ArgsEnvVarReport is a structure that contains all of command line arguments
// and environment variables passed to this process, which can be encoded into
// a JSON representation.
type argsEnvVarReport struct {
  Arguments []string `json:"arguments"`
  EnvVars []string `json:"envvars"`
}

// Report builds a JSON encoded report from the provided arguments and
// environment variables.
func Report(args, envs []string) string {
  report := argsEnvVarReport{ Arguments: args, EnvVars: envs}

  buffer := &bytes.Buffer{}
  if err := json.NewEncoder(buffer).Encode(report); err != nil {
    panic(fmt.Sprintf("%s\n", err.Error()))
  }

  return buffer.String()
}

func main() {
  fmt.Println(Report(os.Args[0:], os.Environ()))

  os.Exit(0)
}
