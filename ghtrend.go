package main

import (
        "flag"
        "fmt"
        "os"
        "github.com/PuerkitoBio/goquery"
)

type RepoInf struct {
        RepositoryName string
        Description    string
}

const (
        TREND_MAX_NUM = 25
)

var repoInf []RepoInf
var baseUrl string = "https://github.com/trending"


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
        n := getNum(*num)
        url := getUrl(*lang)

        repoInf = getMemory(n)

        getPage(url, n)

        showResult()
}

func getUrl(lang string) string {
        if lang == "" {
                return baseUrl
        } else {
                return baseUrl + "?l=" + lang
        }
}

func getNum(num int) int {
        if num > TREND_MAX_NUM {
                num = TREND_MAX_NUM
        }
        return num
}

func getMemory(num int) []RepoInf {
        return make([]RepoInf, num)
}

func getPage(url string, num int) {
        doc, _ := goquery.NewDocument(url)
        doc.Find(".leaderboard-list-content").Each(func(i int, s *goquery.Selection) {
                // fmt.Println(s.Find(".owner-name").Text())
                // fmt.Println(s.Find("span[class='owner-name']").Text())
                // fmt.Println(s.Find("strong").Text())
                if i < num {
                        // fmt.Println(s.Find("a[class='repository-name']").Text())
                        repoInf[i].RepositoryName = s.Find("a[class='repository-name']").Text()
                        repoInf[i].Description = s.Find("p[class='repo-leaderboard-description']").Text()
                }

        })
}

func showResult() {
        fmt.Println("Trending " + *lang + " repositories on GitHub today")
        line := ""
        for i := 0; i < 56; i++ {
                line += "-"
        }
        fmt.Println(line)

        spaces := ""
        for i, rp := range repoInf {
                fmt.Println(fmt.Sprint(i + 1) + ": " + rp.RepositoryName)

                if (i + 1) >= 10 {
                        spaces = "    "
                } else {
                        spaces = "   "
                }
                if *desc {
                        fmt.Println(spaces + rp.Description)
                }
        }
}
