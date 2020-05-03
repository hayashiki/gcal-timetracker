#!/bin/sh

set -o errexit
set -o nounset

archive="TimeTracker-${VERSION}.alfredworkflow"

GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o ".workflow/tt"

(
  cd ./workflow || exit
  envsubst >.info.plist <.info.plist.template
  zip -r "../${archive}" ./*
  zip -d "../${archive}" info.plist.template
)

echo ""
echo "Build completed: \"${archive}\""
