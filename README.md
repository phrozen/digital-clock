![icon](https://github.com/user-attachments/assets/7049cdfe-7f27-4d28-8ebb-a2718305e061)

# digital-clock
A cross platform digital clock widget overlay written in Go. It stays always on top of other windows in your desktop and can be customized to show different formats and styles.

> Although a very simple app, the motivation was to have a cross platform customizable clock that could be used as a status bar widget, with UTC support and that **shows seconds** (*you would be surprised how few apps out there support this... if any*). After some hours of scouring the web for an app, it was clear that there is none that fits my needs, so I decided to write one.

It uses [ebiten](https://github.com/hajimehoshi/ebiten), a lightweight and portable Go library for real-time, high-performance 2D games, widgets and visualizations for showing the clock on screen. 

It also uses [systray](https://github.com/fyne-io/systray), a cross platform library for system tray apps written in Go, which is used to create the icon and menu.

## Configuration

Default configuration file will be created in `./config.toml` on first run. If a different configuration file needs to be used, pass the `-config` flag with the path of the configuration file. On parsing errors, default configuration is loaded, if that happens, check your file.

```toml
# Default Configuration

Seconds = true
Timezone = false
UTC = false
Hours24 = false
FontSize = 64.0

[FontColor]
  R = 255
  G = 255
  B = 255
  A = 255

[Background]
  R = 0
  G = 0
  B = 0
  A = 128
```

When changing the configuration file, click `Reload` from the system tray icon context menu to reload it and show the changes, no need to restart.

## Building from source

### Windows

Youcan just `go build` to get the application, but ideally you want it to have an icon and not show a console window. For that, follow the steps below:

1. Install Go and a C compiler toolchain (CGO is needed for systray)
2. Install rsrc: `go install github.com/akavel/rsrc@latest`
3. Generate `.syso` resource file: `rsrc -arch amd64 -ico tray\icon.ico`
4. Build with GUI flags: `go build -ldflags="-H windowsgui"`

This will get you the binary with an icon and it does not show a window on the taskbar when run.

### Mac OSX

*Coming Soon*

### Linux

To build on Ubuntu/Linux, install the following dependencies for the `hajimehoshi/ebiten` package:

```bash
sudo apt install libc6-dev libgl1-mesa-dev libxcursor-dev libxi-dev libxinerama-dev libxrandr-dev libxxf86vm-dev libasound2-dev pkg-config
```

For more information, refer to the official [hajimehoshi/ebiten documentation](https://ebitengine.org/en/documents/install.html?os=linux).

Install Go dependencies:

```bash
go mod tidy
go mod vendor
```

Build the binary:

```bash
go build -o digital-clock
```

## Usage

Simply run the binary to display the clock on your screen. The clock will remain on top of other windows and can be moved by dragging.

## Credits
All graphic resources are free for personal use and are embedded with the binary when compiled:
- **Font:** digital-7.mono by Sizenko Alexander - [Style-7](http://www.styleseven.com)
- **Icon:** Digital-clock icons created by Freepik - [Flatico](https://www.flaticon.com/free-icons/digital-clock)
