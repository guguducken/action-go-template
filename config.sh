#!/usr/bin/env bash

set -e;

BEFORE_OWNER="guguducken";
BEFORE_REPOSITORY="action-go-template";

OWNER="$1";
REPOSITORY="$2";

FILES=($(grep -r ${BEFORE_OWNER} | grep -Ev 'idea|./.git|config.sh|Makefile' | cut -d ":" -f 1));

for file in "${FILES[@]}"; do
  sed -i "" "s|${BEFORE_OWNER}|${OWNER}|g" "$file";
done

FILES=($(grep -r ${BEFORE_REPOSITORY} | grep -Ev 'idea|./.git|config.sh|Makefile' | cut -d ":" -f 1));

for file in "${FILES[@]}"; do
  sed -i "" "s|${BEFORE_REPOSITORY}|${REPOSITORY}|g" "$file";
done
