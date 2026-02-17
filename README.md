# Year Progress Plugin

A colorful command-line tool and Oh My Zsh plugin that displays the current year's progress as an ASCII progress bar.

## Supported Platforms

| OS | Architecture | Supported |
|----|--------------|-----------|
| Linux | amd64 (x86_64) | ✅ |
| Linux | 386 (i386) | ✅ |
| Linux | ARM | ✅ |
| macOS | amd64 (Intel) | ✅ |
| macOS | arm64 (Apple Silicon) | ✅ |
| FreeBSD | amd64 | ✅ |
| OpenBSD | amd64 | ✅ |

## Features

- 🎨 Random colored progress bar
- ⚙️ Customizable colors via JSON config
- 🖥️ Works as CLI or Oh My Zsh plugin
- 📁 XDG_CONFIG_HOME support
- 🔧 Multiple CLI flags available

## Quick Install (One-Liner)

### With Go (Recommended)
```bash
go install github.com/haroldelopez/year-progress-plugin@latest
```

### With curl (Linux/macOS)
```bash
curl -sSL https://raw.githubusercontent.com/haroldelopez/year-progress-plugin/master/install.sh | bash
```

## Installation

### Package Managers

#### Homebrew (macOS / Linux)

```bash
# Coming soon - formula not yet submitted
# Once published:
brew install year-progress
```
go install github.com/haroldelopez/year-progress-plugin@latest
```

#### Manual Download

Download the latest binary from the [releases](https://github.com/haroldelopez/year-progress-plugin/releases):

```bash
# Linux amd64
curl -L https://github.com/haroldelopez/year-progress-plugin/releases/latest/download/year-progress_linux_amd64 -o year-progress
chmod +x year-progress
./year-progress

# Linux ARM
curl -L https://github.com/haroldelopez/year-progress-plugin/releases/latest/download/year_progress_linux_arm -o year-progress

# macOS Intel
curl -L https://github.com/haroldelopez/year-progress-plugin/releases/latest/download/year_progress_mac_amd64 -o year-progress

# macOS Apple Silicon
curl -L https://github.com/haroldelopez/year-progress-plugin/releases/latest/download/year_progress_mac_arm64 -o year-progress
```

### As Oh My Zsh Plugin
>>>>>>> feat/one-liner-install

### Package Managers

#### Homebrew (macOS / Linux)

```bash
# Coming soon - formula not yet submitted
# Once published:
brew install year-progress
```

#### Go Install

```bash
go install github.com/haroldelopez/year-progress-plugin@latest
```

#### Manual Download

Download the latest binary from the [releases](https://github.com/haroldelopez/year-progress-plugin/releases):

```bash
# Linux amd64
curl -L https://github.com/haroldelopez/year-progress-plugin/releases/latest/download/year-progress_linux_amd64 -o year-progress
chmod +x year-progress
./year-progress

# Linux ARM
curl -L https://github.com/haroldelopez/year-progress-plugin/releases/latest/download/year_progress_linux_arm -o year-progress

# macOS Intel
curl -L https://github.com/haroldelopez/year-progress-plugin/releases/latest/download/year_progress_mac_amd64 -o year-progress

# macOS Apple Silicon
curl -L https://github.com/haroldelopez/year-progress-plugin/releases/latest/download/year_progress_mac_arm64 -o year-progress
```

### As Oh My Zsh Plugin

1. Clone this repository into your Oh My Zsh custom plugins directory:

   ```bash
   # Using HTTPS (recommended)
   git clone https://github.com/haroldelopez/year-progress-plugin.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/year-progress
   
   # Or using SSH
   git clone git@github.com:haroldelopez/year-progress-plugin.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/year-progress
   ```

2. Verify the structure (should have `year_progress.plugin.zsh` and `bin/` folder):
   ```bash
   ls ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/year-progress
   ```

3. Add the plugin to your Oh My Zsh plugins list in `~/.zshrc`:

   ```bash
   plugins=(... year-progress)
   ```

3. Reload your shell configuration:

   ```bash
   source ~/.zshrc
   ```

## Usage

### CLI

```bash
# Show year progress
./year-progress

# Show version
./year-progress --version
./year-progress -v

# Force color output (useful in scripts)
./year-progress --force-color

# Custom config file
./year-progress --config /path/to/colors.json
```

### Oh My Zsh Plugin

Once installed, the Year Progress bar will automatically appear each time you open a new terminal window.

To manually display the Year Progress bar at any time:

```bash
year_progress
```

## Configuration

### Custom Colors

You can customize the colors by creating a JSON configuration file:

1. Create `.year_progress_colors.json` in your home directory or use XDG_CONFIG_HOME:

   ```bash
   # Option 1: Home directory
   nano ~/.year_progress_colors.json

   # Option 2: XDG_CONFIG_HOME
   mkdir -p ~/.config/year-progress
   nano ~/.config/year-progress/.year_progress_colors.json
   ```

2. Add your custom colors:

   ```json
   {
     "Red": "\u001b[31m",
     "Green": "\u001b[32m",
     "Blue": "\u001b[34m",
     "Yellow": "\u001b[33m",
     "Cyan": "\u001b[36m",
     "Magenta": "\u001b[35m",
     "Reset": "\u001b[0m"
   }
   ```

### Supported Color Codes

- Standard ANSI: `\033[30m` - `\033[37m` (colors 0-7)
- Bright colors: `\033[90m` - `\033[97m`
- 256-color mode: `\033[38;5;Nm` (N = 0-255)
- RGB: `\033[38;2;R;G;Bm`

## CLI Options

| Flag | Description |
|------|-------------|
| `-v`, `--version` | Show version information |
| `--config <path>` | Path to custom color config file |
| `--force-color` | Force color output (useful in scripts) |

## Building from Source

Requires Go 1.21+:

```bash
git clone git@github.com:haroldelopez/year-progress-plugin.git
cd year-progress-plugin
go build -ldflags "-X main.Version=1.0.0" -o year-progress .
```

## Running Tests

```bash
go test -v ./...
```

## Troubleshooting

### Permission Denied

If you encounter a "permission denied" error:

```bash
chmod +x ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/year-progress/year_progress.plugin.zsh
```

### Colors Not Showing

Use the `--force-color` flag to force color output:

```bash
./year-progress --force-color
```

## Updating

### CLI

Download the latest release from the [releases page](https://github.com/haroldelopez/year-progress-plugin/releases).

### Plugin

```bash
cd ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/year-progress
git pull
source ~/.zshrc
```

## License

MIT License - see [LICENSE](LICENSE) for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.
