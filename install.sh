#!/bin/bash
# Year Progress Plugin - One-line installer

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

# Map OS
case "$OS" in
    darwin) OS="mac" ;;
    linux|freebsd|openbsd) ;;
    *) echo "Unsupported OS: $OS" && exit 1 ;;
esac

# Build binary name
BINARY_NAME="year-progress_${OS}_${ARCH}"

# Get latest release URL
RELEASE_URL="https://api.github.com/repos/haroldelopez/year-progress-plugin/releases/latest"

# Detect if curl or wget is available
if command -v curl &> /dev/null; then
    DOWNLOAD_CMD="curl -sSL"
elif command -v wget &> /dev/null; then
    DOWNLOAD_CMD="wget -qO-"
else
    echo "Error: curl or wget required"
    exit 1
fi

# Get download URL
DOWNLOAD_URL=$($DOWNLOAD_CMD "$RELEASE_URL" | grep -o "https.*${BINARY_NAME}" | head -1)

if [ -z "$DOWNLOAD_URL" ]; then
    echo "Error: No binary found for $OS/$ARCH"
    echo "Please build from source: go install github.com/haroldelopez/year-progress-plugin@latest"
    exit 1
fi

# Install to ~/.local/bin or /usr/local/bin
if [ -w "$HOME/.local/bin" ]; then
    INSTALL_DIR="$HOME/.local/bin"
elif [ -w "/usr/local/bin" ]; then
    INSTALL_DIR="/usr/local/bin"
else
    INSTALL_DIR="$HOME/.local/bin"
fi

# Download and install
echo "Installing year-progress for $OS/$ARCH..."
$DOWNLOAD_URL "$DOWNLOAD_URL" -o "$INSTALL_DIR/year-progress"
chmod +x "$INSTALL_DIR/year-progress"

echo "✅ Installed to $INSTALL_DIR/year-progress"
echo "Run 'year-progress' to start!"
