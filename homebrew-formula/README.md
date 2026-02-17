# year-progress

A colorful CLI tool that displays the current year's progress as an ASCII progress bar.

## Installation

```bash
# Add the tap
brew tap haroldelopez/year-progress

# Install
brew install year-progress
```

## Usage

```bash
# Show year progress
year-progress

# Show version
year-progress --version

# Use custom colors
year-progress --config ~/my-colors.json
```

## Options

- `--version` - Show version
- `--config` - Path to custom color config file
- `--force-color` - Force color output

## Custom Colors

Create a `~/.year_progress_colors.json` file with your colors:

```json
{
  "#FF6B6B": "#FF6B6B",
  "#4ECDC4": "#4ECDC4",
  "Reset": "\u001b[0m"
}
```

## License

MIT - See LICENSE file for details.
