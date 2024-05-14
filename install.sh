#!/bin/bash

set -e

# Function to install dependencies for Debian/Ubuntu
install_deps_debian() {
  sudo apt-get update
  sudo apt-get install -y libvips-dev golang-go
}

# Function to install dependencies for macOS
install_deps_macos() {
  brew update
  brew install vips go
}

# Function to install dependencies for Windows (using chocolatey)
install_deps_windows() {
  choco install -y vips golang
}

# Detect the operating system and install dependencies
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  . /etc/os-release
  if [[ "$ID" == "ubuntu" || "$ID" == "debian" ]]; then
    install_deps_debian
  else
    echo "Unsupported Linux distribution. Please install libvips-dev and golang-go manually."
    exit 1
  fi
elif [[ "$OSTYPE" == "darwin"* ]]; then
  install_deps_macos
elif [[ "$OSTYPE" == "msys" ]]; then
  install_deps_windows
else
  echo "Unsupported OS. Please install libvips and Go manually."
  exit 1
fi

# Set PKG_CONFIG_PATH for pkg-config to find vips
if [[ "$OSTYPE" == "linux-gnu"* || "$OSTYPE" == "darwin"* ]]; then
  export PKG_CONFIG_PATH=/usr/local/lib/pkgconfig:/usr/lib/pkgconfig
  echo 'export PKG_CONFIG_PATH=/usr/local/lib/pkgconfig:/usr/lib/pkgconfig' >> ~/.bashrc
  # shellcheck disable=SC1090
  source ~/.bashrc
elif [[ "$OSTYPE" == "msys" ]]; then
  echo "Please add the directory containing 'vips.pc' to the PKG_CONFIG_PATH environment variable manually."
fi

# Install Go packages
go get -u github.com/h2non/bimg
go get -u github.com/spf13/cobra

# Build the application
go build -o image-processor app/main.go

echo "Setup and build completed successfully."
