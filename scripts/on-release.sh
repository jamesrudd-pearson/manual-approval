#!/bin/bash

RELEASE_VERSION=$1

if [ -z "$RELEASE_VERSION" ]; then
  echo "RELEASE_VERSION is not set"
  exit 1
fi

MAJOR_VERSION=$(echo $RELEASE_VERSION | cut -d. -f1)

sed -i "s|\(image: docker://ghcr.io/jamesrudd-pearson/manual-approval:\)[0-9]\+\.[0-9]\+\.[0-9]\+|\1$RELEASE_VERSION|g" action.yaml

git add action.yaml
git commit -m "Release $RELEASE_VERSION [skip ci]"
git push origin main
git tag -f v$RELEASE_VERSION && git push -f origin v$RELEASE_VERSION
git tag -f v$MAJOR_VERSION && git push -f origin v$MAJOR_VERSION
