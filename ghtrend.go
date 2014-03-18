package main

import (
        "flag"
        "fmt"
        "os"
)

var (
        lang = flag.String("l", "all", "Select language")
        desc = flag.Bool("d", false, "Show description")
        num  = flag.Int("n", 10, "Limit numbers")
        help = flag.Bool("h", false, "Show help message")
)
func main() {
        flag.Usage = func() {
                fmt.Fprint(os.Stderr, `
usage: ghtrend <command> [options] <args>

optional arguments:
  -l    Select language.
  -d    Show description.
  -n    Limit numbers.
  -h    Show help message.
`)
        }
        flag.Parse()

        if *help {
                flag.Usage()
                os.Exit(0)
        }

}
