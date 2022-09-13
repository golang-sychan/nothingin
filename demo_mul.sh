#!/usr/bin/env bash

set -e
set -o pipefail

function foo() {
  echo "foo-000"
  return 9
}

echo "0000"
echo "hehe---00"  | echo "hehe---999"
echo "9999"
