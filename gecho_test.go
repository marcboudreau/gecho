package main

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestReport(t *testing.T) {
  testcases := []struct{
    args []string
    envs []string
    expected string
  } {
    {
      args: []string{"first"},
      envs: []string{},
      expected: "{\"arguments\":[\"first\"],\"envvars\":[]}\n",
    },
    {
      args: []string{"one", "two", "three"},
      envs: []string{},
      expected: "{\"arguments\":[\"one\",\"two\",\"three\"],\"envvars\":[]}\n",
    },
    {
      args: []string{"gecho", "-c", "12", "--verbose"},
      envs: []string{},
      expected: "{\"arguments\":[\"gecho\",\"-c\",\"12\",\"--verbose\"],\"envvars\":[]}\n",
    },
    {
      args: []string{"binary", "", "\\", "\""},
      envs: []string{},
      expected: "{\"arguments\":[\"binary\",\"\",\"\\\\\",\"\\\"\"],\"envvars\":[]}\n",
    },
    {
      args: []string{"binary", "", "\\", "\""},
      envs: []string{"DEBUG=1"},
      expected: "{\"arguments\":[\"binary\",\"\",\"\\\\\",\"\\\"\"],\"envvars\":[\"DEBUG=1\"]}\n",
    },
    {
      args: []string{"program", "-a", "file.txt"},
      envs: []string{"HOME=/home/user", "SHELL=/usr/bin/bash", "QUOTE=\""},
      expected: "{\"arguments\":[\"program\",\"-a\",\"file.txt\"],\"envvars\":[\"HOME=/home/user\",\"SHELL=/usr/bin/bash\",\"QUOTE=\\\"\"]}\n",
    },
  }

  for _, testcase := range testcases {
    result := Report(testcase.args, testcase.envs)

    assert.Equal(t, result, testcase.expected)
  }
}
