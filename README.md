# taildiff

Monitor changes of a supplied shell command.

[![demo](https://asciinema.org/a/OoYxvY16YlBMBn4Rv9C8gUarh.svg)](https://asciinema.org/a/OoYxvY16YlBMBn4Rv9C8gUarh?autoplay=1)

## Quick intro

taildiff lets you monitor/record the changes of a shell command.

Supply the shell command with `-c '<command>'`, e.g. `-c 'grep processes /proc/stat'`.

The default update interval is 1 second, change it with `-n <duration>`, e.g. `-n 0.3s`

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
