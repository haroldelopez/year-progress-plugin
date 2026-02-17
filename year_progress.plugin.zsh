#!/usr/bin/env zsh

# Get the directory where this plugin is located
PLUGIN_DIR="${0:a:h}"

# Determine the correct binary based on the system
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  if [[ "$(uname -m)" == "x86_64" ]]; then
    BINARY="year_progress_linux_amd64"
  else
    BINARY="year_progress_linux_386"
  fi
elif [[ "$OSTYPE" == "darwin"* ]]; then
  if [[ "$(uname -m)" == "arm64" ]]; then
    BINARY="year_progress_mac_arm64"
  else
    BINARY="year_progress_mac_amd64"
  fi
else
  echo "Unsupported operating system"
  return 1
fi

# Function to run the year progress binary
function year_progress() {
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
