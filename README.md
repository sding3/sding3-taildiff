# taildiff

Monitor changes of a supplied shell command.

[![demo](https://asciinema.org/a/OoYxvY16YlBMBn4Rv9C8gUarh.svg)](https://asciinema.org/a/OoYxvY16YlBMBn4Rv9C8gUarh?autoplay=1)

## Quick intro

taildiff lets you monitor/record the changes of a shell command.

`-c '<command>'` supplies the shell command, e.g. `-c 'grep processes /proc/stat'`.

`-n <duration>` is optional. It supplies update interval, e.g. `-n 0.3s`. Default is 1s.

Other options are available too:

```
$ taildiff -h
Usage of taildiff:
  -c string
        [required] shell command
  -e    exit on command error. (default false)
  -n duration
        update interval (default 1s)
```

## Installation

### go get

```
go get github.com/sding3/taildifff
```

### Compile from source

```
git clone https://github.com/sding3/taildiff.git
cd taildiff
make build
sudo make install # installs to /usr/loca/bin/taildiff
```
