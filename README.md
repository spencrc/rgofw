# RGoFW
Go bindings for [RGFW](https://github.com/ColleagueRiley/RGFW), "a cross platform lightweight single-header simple-to-use window abstraction library for creating graphical programs or libraries." These bindings target RGFW 1.8.1, and will be updated when a new major version releases. 

MacOS is currently **unsupported** by RGoFW. If you would like to add support, please create a pull request!

## Installation
```bash
go get github.com/spencrc/RGoFW
```

## Dependencies
On Linux, you require the following packages:
- libX11
- libXrandr

For Wayland specifically, you require the following packages:
- libwayland
- libxkbcommon

## Credits
RGFW is licensed under the zlib license by Riley Mabb.
