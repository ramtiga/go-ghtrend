# go-ghtrend

Get Trending repositories on Github written by Go.

## Installation

    go get github.com/ramtiga/go-ghtrend

## Usage

    $ ghtrend -h

    usage: ghtrend <command> [options] <args>
    
    optional arguments:
      -l    Select language.
      -d    Show description.
      -n    Limit numbers.
      -b    Show repository on browser.
      -h    Show help message.
      -v    Show version.

## Example

    $ ghtrend -l go -n 5 -d
    Trending go repositories on GitHub today
    --------------------------------------------------------
    1: nytlabs/streamtools
       tools for working with streams of data
    2: oysterbooks/halfshell
       A proxy server for processing images on the fly.
    3: dotcloud/docker
       Docker - the open-source application container engine
    4: mitchellh/go-mruby
       Go (golang) bindings to mruby.
    5: drone/drone
       Drone is a Continuous Integration platform built on Docker)

## License

MIT

## Author

ramtiga

