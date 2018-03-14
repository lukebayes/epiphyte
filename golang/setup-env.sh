#!/usr/bin/env bash

# Run this in terminal with `source setup-env.sh.
if [ -n "${ZSH_VERSION}" ]; then
  BASEDIR="$( cd $( dirname "${(%):-%N}" ) && pwd )"
elif [ -n "${BASH_VERSION}" ]; then
  if [[ "$(basename -- "$0")" == "setup-env.sh" ]]; then
    echo "Don't run $0, source it (see README.md)" >&2
    exit 1
  fi
  BASEDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
else
  echo "Unsupported shell, use bash or zsh."
  exit 2
fi

BASEDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
source "${BASEDIR}/script/utils.sh"

# Add script to the path.
add_to_path "${BASEDIR}/script"
echo "Updated PATH with ${BASEDIR}/script"

# Add go lang bin to the path.
add_to_path "${BASEDIR}/lib/go-1.10/bin"
echo "Updated PATH with ${BASEDIR}/lib/go-1.10/bin"

# Add project bin to the path.
add_to_path "${BASEDIR}/bin/"
echo "Updated PATH with ${BASEDIR}/bin"

# Add vendor bin to the path.
add_to_path "${BASEDIR}/vendor/bin/"
echo "Updated PATH with ${BASEDIR}/vendor/bin"

# Add Google depot_tools to path.
add_to_path "${BASEDIR}/lib/depot_tools"
echo "Updated PATH with ${BASEDIR}/lib/depot_tools"

# Set the custom GOROOT value to the local Go installation.
export GOROOT="${BASEDIR}/lib/go-1.10"
echo "Set GOROOT to ${BASEDIR}/lib/go-1.10"

# Set the gopath for this project
export GOPATH="${BASEDIR}:${BASEDIR}/vendor"
echo "Set GOPATH to $GOPATH"

# Set CGO_LDFLAGS so that CGO can access Skia libraries
export CGO_LDFLAGS="-L ${BASEDIR}/lib/skia/out/Shared -Wl -rpath ${BASEDIR}/lib/skia/out/Shared -lskia"
echo "Set CGO_LDFLAGS=${CGO_LDFLAGS}"

# Set CGO_CFLAGS so that CGO can access Skia libraries
export CGO_CFLAGS="-I${BASEDIR}/lib/skia/include/c"
echo "Set CGO_CFLAGS=${CGO_CFLAGS}"

