#!/usr/bin/env zsh

# Get the directory where this plugin is located
PLUGIN_DIR="${0:a:h}"

# Debug: uncomment to debug
# echo "DEBUG: PLUGIN_DIR=$PLUGIN_DIR"
# echo "DEBUG: OSTYPE=$OSTYPE"
# echo "DEBUG: ARCH=$(uname -m)"

# Determine the correct binary based on the system
if [[ "$OSTYPE" == "darwin"* ]]; then
  ARCH=$(uname -m)
  if [[ "$ARCH" == "arm64" ]]; then
    BINARY="year_progress_mac_arm64"
  elif [[ "$ARCH" == "x86_64" ]]; then
    BINARY="year_progress_mac_amd64"
  else
    echo "Unsupported Apple Silicon architecture: $ARCH"
    return 1
  fi
elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
  ARCH=$(uname -m)
  if [[ "$ARCH" == "x86_64" ]]; then
    BINARY="year_progress_linux_amd64"
  elif [[ "$ARCH" == "i386" ]] || [[ "$ARCH" == "i686" ]]; then
    BINARY="year_progress_linux_386"
  elif [[ "$ARCH" == "aarch64" ]] || [[ "$ARCH" == "arm64" ]]; then
    BINARY="year_progress_linux_arm"
  else
    echo "Unsupported Linux architecture: $ARCH"
    return 1
  fi
else
  echo "Unsupported operating system: $OSTYPE"
  return 1
fi

# Function to run the year progress binary
function year_progress() {
  if [[ ! -f "$PLUGIN_DIR/bin/$BINARY" ]]; then
    echo "Error: Binary not found: $PLUGIN_DIR/bin/$BINARY"
    echo "Available files:"
    ls -la "$PLUGIN_DIR/bin/" 2>/dev/null || echo "No bin directory found"
    return 1
  fi
  "$PLUGIN_DIR/bin/$BINARY"
}

# Function to check if the plugin should run
function should_run_year_progress() {
  local last_run_file="${HOME}/.year_progress_last_run"
  local current_date=$(date +%Y%m%d)

  if [[ ! -f "$last_run_file" ]] || [[ $(cat "$last_run_file") != "$current_date" ]]; then
    echo "$current_date" > "$last_run_file"
    return 0
  else
    return 1
  fi
}

# Run year_progress when a new shell is opened, but only once per day
if should_run_year_progress; then
  year_progress
fi
