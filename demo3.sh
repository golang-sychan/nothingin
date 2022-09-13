#!/usr/bin/env bash
set -e

function foo() {
  echo "sub-000"
  ls *.notexist | echo "hello"
  echo "sub-999"
  return 9
}

echo "0000"
foo
echo "9999"

