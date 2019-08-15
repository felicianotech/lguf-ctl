# LG UltraFine Go Library [![CircleCI Build Status](https://circleci.com/gh/felicianotech/go-lguf.svg?style=shield)](https://circleci.com/gh/felicianotech/go-lguf) [![GitHub License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/felicianotech/go-lguf/master/LICENSE)

`go-lguf` is a Go library that interfaces with the LG UltraFine 4K monitor in order to adjust brightness from Linux.
This simple library is needed because this monitor was designed specifically for Apple computers and thus has no physical buttons.
Without the built-in features of macOS, adjusting brightness on this monitor wasn't possible.


## Requirements

- Ubuntu 18.04 "Bionic", Ubuntu 19.04 "Disco" - This project supports 64-bit desktop installations of Ubuntu 18.04 "Bionic" 
and Ubuntu 19.04 "Disco".
For 32-bit support or other distros, please open a Pull Request or Issue.


## Installation

`zoom-mgr` is a Bash script that can be run from anywhere in your home 
directory.
Here's a typical install:

```
cd ~
curl -O https://raw.githubusercontent.com/felicianotech/zoom-mgr/master/zoom-mgr.sh
chmod +x ~/zoom-mgr.sh
```


## Usage

```
~/zoom-mgr.sh install
```

Installs the Zoom Linux Client on your machine.

```
~/zoom-mgr.sh update
```

If a newer version of Zoom is available, download and install it.


## Development

Instructions coming in the future.


## License

The source code for `zoom-mgr` is licensed under the MIT license.
This repo's license can be found [here](./LICENSE).
Zoom itself is a proprietary product by [Zoom Video Communications, Inc](https://zoom.us/).
This project is not affiliated with nor endorsed by them.
