#!/bin/bash
# Year Progress Plugin - One-line installer
# Usage: curl -sSL https://raw.githubusercontent.com/haroldelopez/year-progress-plugin/master/install.sh | bash

set -e

# Detect OS and architecture
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

# Map architecture
case "$ARCH" in
    x86_64) ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
    arm*) ARCH="arm" ;;
    i386|i686) ARCH="386" ;;
esac

# Map OS to match release naming
case "$OS" in
    darwin) OS="mac" ;;
    linux|freebsd|openbsd) ;;
    *) echo "Unsupported OS: $OS" && exit 1 ;;
esac

# Build binary name
BINARY_NAME="year-progress_${OS}_${ARCH}"
BASE_URL="https://github.com/haroldelopez/year-progress-plugin/releases/download/v1.0.0"
DOWNLOAD_URL="${BASE_URL}/${BINARY_NAME}"

# Install to ~/.local/bin or /usr/local/bin
if [ -w "$HOME/.local/bin" ]; then
    INSTALL_DIR="$HOME/.local/bin"
elif [ -w "/usr/local/bin" ]; then
    INSTALL_DIR="/usr/local/bin"
else
    INSTALL_DIR="$HOME/.local/bin"
fi

# Ensure install dir exists
mkdir -p "$INSTALL_DIR"

# Download and install
echo "Installing year-progress for $OS/$ARCH..."

if command -v curl &> /dev/null; then
    if curl -sSL --fail -o "$INSTALL_DIR/year-progress" "$DOWNLOAD_URL"; then
        chmod +x "$INSTALL_DIR/year-progress"
        echo "✅ Installed to $INSTALL_DIR/year-progress"
        echo ""
        echo "Run 'year-progress' to start!"
    else
        rm -f "$INSTALL_DIR/year-progress"
        echo "❌ Binary not found for $OS/$ARCH"
        echo "Building from source instead..."
        if command -v go &> /dev/null; then
            go install github.com/haroldelopez/year-progress-plugin@latest
            echo "✅ Installed via Go"
        else
            echo "❌ Go not installed. Please install from: https://go.dev/dl/"
            exit 1
        fi
    fi
elif command -v wget &> /dev/null; then
    if wget -qO "$INSTALL_DIR/year-progress" "$DOWNLOAD_URL"; then
        chmod +x "$INSTALL_DIR/year-progress"
        echo "✅ Installed to $INSTALL_DIR/year-progress"
        echo "Run 'year-progress' to start!"
    else
        rm -f "$INSTALL_DIR/year-progress"
        echo "❌ Failed to download"
        exit 1
    fi
else
    echo "Error: curl or wget required"
    exit 1
fi
