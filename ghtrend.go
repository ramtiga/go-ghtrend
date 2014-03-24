package main

import (
        "flag"
        "fmt"
        "os"
        "github.com/PuerkitoBio/goquery"
        "os/exec"
        "runtime"
)

type RepoInf struct {
        RepositoryName string
        Description    string
        RepoUrl        string
}

const (
        TREND_MAX_NUM = 25
        VERSION = "0.0.1"
)

var repoInf []RepoInf
var baseUrl string = "https://github.com/trending"


var (
        lang    = flag.String("l", "all", "Select language")
        desc    = flag.Bool("d", false, "Show description")
        num     = flag.Int("n", 10, "Limit numbers")
        brows   = flag.Int("b", 0, "Show repository on browser")
        help    = flag.Bool("h", false, "Show help message")
        version = flag.Bool("v", false, "Show version")
)

func main() {
        flag.Usage = func() {
                fmt.Fprint(os.Stderr, `
usage: ghtrend <command> [options] <args>

optional arguments:
  -l    Select language.
  -d    Show description.
  -n    Limit numbers.
  -b    Show repository on browser.
  -h    Show help message.
  -v    Show version.
`)
        }
        flag.Parse()

        if *version {
                showVersion()
                os.Exit(0)
        }
        if *help {
                flag.Usage()
                os.Exit(0)
        }
        
        n := getNum(*num)
        url := getGithubUrl(*lang)

        repoInf = getMemory(n)

        getPage(url, n)

        if *brows > 0 && *brows <= 25 {
                browsUrl := getBrowsUrl(*brows)
                openBrowser(browsUrl)
                os.Exit(0)
        }
        showResult()
}

func getGithubUrl(lang string) string {
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
                if i < num {
                        repoInf[i].RepositoryName = s.Find("a[class='repository-name']").Text()
                        repoInf[i].Description = s.Find("p[class='repo-leaderboard-description']").Text()
                        repoInf[i].RepoUrl = s.Find("a[class='repository-name']").Text()
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

func getBrowsUrl(idx int) string {
        return "https://github.com/" + repoInf[idx - 1].RepoUrl
}

func openBrowser(url string) {
        os := runtime.GOOS
        switch {
        case os == "windows":
                exec.Command("cmd", "/c", "start", url).Run()
        case os == "darwin":
                exec.Command("open", url).Run()
        case os == "windows":
                exec.Command("xdg-open", url).Run()
        }
}

func showVersion() {
        fmt.Printf("ghtrend Ver %s\n", VERSION)
}
