# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-02-15

### Added
- Version flag (`-v`, `--version`)
- Config file support via `--config` flag
- XDG_CONFIG_HOME support for config file location
- Force color mode (`--force-color`)
- Unit tests for core functions
- GoReleaser configuration for automated releases
- CHANGELOG.md

### Fixed
- Config file path now looks in home directory first
- Color rendering in non-TTY environments

### Changed
- Refactored main.go with better flag handling
- Added more default colors (Cyan, Magenta)
- Error handling improvements

---

## [0.0.0] - YYYY-MM-DD (Initial Release)
- Initial release with basic year progress bar
