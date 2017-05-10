package main

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestReport(t *testing.T) {
  testcases := []struct{
    args []string
    expected string
  } {
    {args: []string{"first"}, expected: "{\"arguments\":[\"first\"]}\n"},
    {args: []string{"one", "two", "three"}, expected: "{\"arguments\":[\"one\",\"two\",\"three\"]}\n"},
    {args: []string{"gecho", "-c", "12", "--verbose"}, expected: "{\"arguments\":[\"gecho\",\"-c\",\"12\",\"--verbose\"]}\n"},
    {args: []string{"binary", "", "\\", "\""}, expected: "{\"arguments\":[\"binary\",\"\",\"\\\\\",\"\\\"\"]}\n"},
  }

  for _, testcase := range testcases {
    result := Report(testcase.args)

    assert.Equal(t, result, testcase.expected)
  }
}
