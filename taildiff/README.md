# taildiff

Monitor changes of a supplied shell command's output.

[![demo](https://asciinema.org/a/304828.svg)](https://asciinema.org/a/304828?autoplay=1)

## Quick intro

taildiff lets you monitor/record the changes of a shell command.

`-c '<command>'` supplies the shell command, e.g. `-c 'grep processes /proc/stat'`.

Other options are available too:

```
$ taildiff -h
taildiff monitors changes of a supplied shell command's stdout.

Usage of ./taildiff:
  -c string
    	[required] shell command
  -e	exit on command error
  -n duration
    	update interval (default 1s)
  -no-newline
    	don't print the extra newline character after each output
  -time-stamp
    	prefix timstamp to each output line
```

Examples:

```
taildiff -c 'grep processes /proc/stat' -no-newline -time-stamp
taildiff -c 'grep processes /proc/stat'
taildiff -c 'netstat -tupn' -time-stamp
```

## Installation

### Compile from source

First download and install [go](https://golang.org/doc/install) if you don't have it, then perform the following.

```
git clone https://github.com/sding3/taildiff.git
cd taildiff
make build
sudo make install # installs to /usr/loca/bin/taildiff
```
