# go-ghtrend

Get Trending repositories on Github written by Go.

## Installation

    go get github.com/ramtiga/go-ghtrend

## Usage

    $ go-ghtrend -h

    usage: ghtrend <command> [options] <args>
    
    optional arguments:
      -l    Select language.
      -d    Show description.
      -n    Limit numbers.
      -b    Show repository on browser.
      -h    Show help message.
      -v    Show version.

## Example

    $ go-ghtrend -l go
    Trending go repositories on GitHub today
    No. Name                                 Star Fork
    --- ------------------------------------ ---- ----
      1 GoogleCloudPlatform/kubernetes         46    4
      2 aybabtme/uniplot                       46
      3 shell909090/goproxy                    44    1
      4 andlabs/ui                             31
      5 dotcloud/docker                        24   13
      6 hellogcc/100-gdb-tips                  24    5
      7 lestrrat/peco                          26
      8 google/cadvisor                        20
      9 docker/libswarm                        18    2
     10 calmh/syncthing                        16    1

## License

MIT

## Author

ramtiga

