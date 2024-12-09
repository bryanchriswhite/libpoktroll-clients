## Poktroll C Clients Library

This repo contains an asynchronous C API to the [Poktroll client packages](https://pkg.go.dev/github.com/pokt-network/poktroll/pkg/client) via [cgo](https://pkg.go.dev/cmd/cgo) wrapper functions.

```bash
# Clone and cd into the repo.
git clone https://github.com/bryanchriswhite/libpoktroll_clients.git --recurse-submodules
cd libpoktroll_clients

# If you cloned but didn't pull the submodules, run:
git submodule update --init --recursive

# (optional) Update protobufs ("pull" from buf.build)
buf export buf.build/pokt-network/poktroll

# Generate protobufs
buf generate

# Make and cd to cmake build directory.
mkdir build
cd build

# Generate build files and build the library..
cmake ..
make

# Run tests (requires running poktroll localnet.
# (see: https://dev.poktroll.com/develop/developer_guide/quickstart#1-launch--inspect-localnet)
ctest --output-on-failure

# Build deb/rpm/tar install packages.
make package  # ALL

## Produces:
#  - build/libpoktroll_clients-<version>-Linux.{sh,tar.gz,tar.Z}
#  - build/libpoktroll_clients-<version>.deb
#  - build/libpoktroll_clients-<version>_amd64.deb
#  - build/libpoktroll_clients-<version>.x86_64.rpm

## OR
cpack -G "TGZ;DEB;RPM"  # All
cpack -G DEB            # Debian
cpack -G RPM            # RHEL/Fedora
cpack -G TGZ            # tar.gz

# Build arch install package (depends on TGZ from cpack).
make pkgbuild

## Produces:
#  - build/PKGBUILD
#  - build/pkg/*

# Install the shared library and headers.

## Arch
sudo pacman -U ./pkg/libpoktroll_clients-0.1.0-1-x86_64.pkg.tar.zst   

## Debian
sudo dpkg -i ./pkg/libpoktroll_clients-0.1.0-Linux.deb

## RHEL/CentOS
sudo rpm -i ./pkg/libpoktroll_clients-0.1.0-Linux.rpm
```