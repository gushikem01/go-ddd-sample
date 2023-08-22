#!/bin/bash
set -e

DIR="$(cd "$(dirname "$0")" && pwd -P)"
packages="$(cd "$DIR/../" && go work edit -json | jq -c -r '[.Use[].DiskPath] | map_values(.)[]')"
echo "$packages" | while read -r endpoint; do
    absolute="$(cd "$DIR/../" && cd "$endpoint" && pwd)"
    echo "$absolute go mod tidy"
    cd "$absolute" && go mod tidy
done
