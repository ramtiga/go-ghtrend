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
        Lang           string
        Star           string
        Fork           string
}

const (
        TREND_MAX_NUM = 25
        REPO_NAME_MAX_LEN = 35
        LANG_LEN_MAX = 12
)

var repoInf []RepoInf
var baseUrl string = "https://github.com/trending"
var repoNameMaxLen int

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
        repoNameMaxLen = REPO_NAME_MAX_LEN
        var repolen int
        doc, _ := goquery.NewDocument(url)
        doc.Find(".leaderboard-list-content").Each(func(i int, s *goquery.Selection) {
                if i < num {
                        repoInf[i].RepositoryName = s.Find("a[class='repository-name']").Text()
                        repoInf[i].Description = s.Find("p[class='repo-leaderboard-description']").Text()
                        repoInf[i].RepoUrl = s.Find("a[class='repository-name']").Text()
                        repoInf[i].Lang = s.Find("span[class='title-meta']").Text()
                        repoInf[i].Star = s.Find(".repo-leaderboard-meta .repo-leaderboard-meta-item .octicon-star").Parent().Text()
                        repoInf[i].Fork = s.Find(".repo-leaderboard-meta .repo-leaderboard-meta-item .octicon-git-branch").Parent().Text()

                        repolen = len(repoInf[i].RepositoryName) 
                        if repolen > repoNameMaxLen {
                                repoNameMaxLen = repolen
                        }
                }

        })
}

func showResult() {
        fmt.Println("Trending " + *lang + " repositories on GitHub today")

        spaces := " "
        line := ""
        for i := 0; i < repoNameMaxLen - 4; i++ {
                spaces += " "
                line += "-"
        }

        title := ""
        lines := ""
        title_starfork := "Star Fork"
        line_starfork := " ---- ----"

        if *lang == "all" {
                title = "No. Name " + spaces + "Lang         " + title_starfork
                lines = "--- -----" + line  + " ------------" + line_starfork
        } else {
                title = "No. Name " + spaces + title_starfork
                lines = "--- -----" + line + line_starfork
        }

        fmt.Println(title)
        fmt.Println(lines)

        for i, rp := range repoInf {
                spaces = "  "
                for j := 0; j < repoNameMaxLen - len(rp.RepositoryName); j++ {
                        spaces += " "
                }
                spaces2 := " "

                if *lang == "all" {
                        for k := 0; k < LANG_LEN_MAX - len(rp.Lang); k++ {
                                spaces2 += " "
                        }
                        fmt.Println(fmt.Sprintf("%3d", i + 1) + " " + rp.RepositoryName + spaces + rp.Lang + spaces2 + fmt.Sprintf("%4s", rp.Star) + " " + fmt.Sprintf("%4s", rp.Fork))
                } else {
                        fmt.Println(fmt.Sprintf("%3d", i + 1) + " " + rp.RepositoryName + spaces + fmt.Sprintf("%4s", rp.Star) + " " + fmt.Sprintf("%4s", rp.Fork))
                }

                if *desc {
                        fmt.Println("    " + rp.Description)
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
        case os == "linux":
                exec.Command("xdg-open", url).Run()
        }
}

