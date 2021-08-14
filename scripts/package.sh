#!/bin/bash

set -e
PRJROOT=$(dirname "${BASH_SOURCE[0]}")/..
cd $PRJROOT

PKGNAME="dnazart"

PLATFORM="$1"
GOOS="$2"
GOARCH="$3"
VERSION=$(git describe --tags --abbrev=0)

echo Packaging $PLATFORM Binary

# Remove previous build directory, if needed.
bdir=$PKGNAME-$VERSION-$GOOS-$GOARCH
rm -rf packages/$bdir && mkdir -p packages/$bdir

# Make the binaries.
GOOS=$GOOS GOARCH=$GOARCH make all

# Copy the executable binaries.
if [ "$GOOS" == "windows" ]; then
	mv tmp/dnazart packages/$bdir/dnazart.exe
else
	mv tmp/dnazart packages/$bdir
fi

# Copy documention and license.
cp README.md packages/$bdir
#cp CHANGELOG.md packages/$bdir
#cp LICENSE packages/$bdir

# Compress the package.
cd packages
if [ "$GOOS" == "linux" ]; then
	tar -zcf $bdir.tar.gz $bdir
else
	zip -r -q $bdir.zip $bdir
fi

# Remove build directory.
rm -rf $bdir
