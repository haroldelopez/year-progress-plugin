# Year Progress Oh My Zsh Plugin

A colorful command-line tool and Oh My Zsh plugin that displays the current year's progress as an ASCII progress bar.

## Installation

1. Clone this repository into your Oh My Zsh custom plugins directory:

   ```
   git clone git@github.com:haroldelopez/year-progress-plugin.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/year-progress
   ```

2. Add the plugin to your Oh My Zsh plugins list in `~/.zshrc`:

   ```
   plugins=(... year-progress)
   ```

   Make sure to add `year-progress` to your existing list of plugins.

3. Reload your shell configuration:

   ```
   source ~/.zshrc
   ```

## Usage

Once installed, the Year Progress bar will automatically appear each time you open a new terminal window.

To manually display the Year Progress bar at any time, simply type:

```
year_progress
```

## Customization

You can customize the colors used in the progress bar by creating a JSON configuration file:

1. Create a hidden file named `.year_progress_colors.json` in your home directory:

   ```
   nano ~/.year_progress_colors.json
   ```

2. Add your custom colors using the following format:

   ```json
   {
     "Red": "\u001b[31m",
     "Green": "\u001b[32m",
     "Blue": "\u001b[34m",
     "Yellow": "\u001b[33m",
     "Reset": "\u001b[0m"
   }
   ```

3. Modify the color codes as desired. The plugin will automatically use these custom colors.

## Troubleshooting

If you encounter a "permission denied" error, make the plugin file executable:

```
chmod +x ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/year-progress/year_progress.plugin.zsh
```

## Updating

To update the plugin to the latest version:

1. Navigate to the plugin directory:

   ```
   cd ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/year-progress
   ```

2. Pull the latest changes:

   ```
   git pull
   ```

3. Restart your terminal or run:

   ```
   source ~/.zshrc
   ```

Enjoy tracking the year's progress in style!