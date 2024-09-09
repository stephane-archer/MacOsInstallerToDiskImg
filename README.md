# MacOsInstallerToDiskImg

`MacOsInstallerToDiskImg` is a command-line utility that converts a macOS installer obtained from the App Store into a bootable disk image (ISO).
This disk image can be used to install macOS on a virtual machine (VM), or for other purposes that require an ISO image.

## Features
- Converts a macOS installer app into a bootable disk image (ISO).
- Simple and fast conversion process.

## Requirements
- Go (for building the project)
- macOS installer from the App Store (e.g., `/Applications/Install macOS Mojave.app/`)

## Installation

### Option 1: Install with `go install`

You can easily install `MacOsInstallerToDiskImg` using Go by running the following command:
```bash
go install https://github.com/stephane-archer/MacOsInstallerToDiskImg@latest
```

This will install the executable in your `$GOPATH/bin` directory, making it available for use system-wide.

### Option 2: Build from Source

If you prefer to build the utility from source, follow these steps:
1. Ensure you have Go installed on your system.
2. Clone or download this repository.
3. Open a terminal and navigate to the project directory.
4. Run the following command to build the binary:
    ```bash
    go build
    ```
    This will generate an executable file named `MacOsInstallerToDiskImg` in the current directory.

## Usage

### Step 1: Download a macOS Installer

Ensure you have downloaded a macOS installer from the App Store.
You can typically find it in the `/Applications` folder.
For example:
```bash
/Applications/Install macOS Mojave.app/
```
### Step 2: Run the Command

To convert the macOS installer into a disk image, run the following command:
```bash
MacOsInstallerToDiskImg /path/to/macos/installer.app /path/to/output.iso
```
For example, to convert the macOS Mojave installer into an ISO file, you would run:
```bash
MacOsInstallerToDiskImg /Applications/Install\ macOS\ Mojave.app/ ~/MacOSMojave.iso
```
This will create an ISO image of the macOS Mojave installer in your home directory with the name `MacOSMojave.iso`.
