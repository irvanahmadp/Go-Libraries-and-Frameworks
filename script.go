package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

//App struct
type App struct {
	LastUpdate         int64      `json:"last_update"`
	RateLimit          int        `json:"rate_limit"`
	RateLimitRemaining int        `json:"rate_limit_remaining"`
	RateLimitReset     int64      `json:"rate_limit_reset"`
	Categories         []Category `json:"categories"`
}

//Category struct
type Category struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	LastUpdate  int64  `json:"last_update"`
}

//ListURL struct
type ListURL struct {
	URLs []string `json:"URLs"`
}

//CategiryData struct
type CategoryData struct {
	Name           string    `json:"name"`
	IndexLastCheck int       `json:"index_last_check"`
	LastUpdate     int64     `json:"last_update"`
	Libraries      []Library `json:"libraries"`
}

//Library struct
type Library struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	RepoURL     string    `json:"html_url"`
	HomePageURL string    `json:"homepage"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	SvnURL      string    `json:"svn_url"`
	Stargazers  int       `json:"stargazers_count"`
	Watchers    int       `json:"watchers_count"`
	Forks       int       `json:"forks_count"`
	OpenIssues  int       `json:"open_issues_count"`
}

func main() {
	appData := App{}
	appFileName := "app.json"
	appFile := readFile(appFileName)
	err := json.Unmarshal(appFile, &appData)
	categories := appData.Categories
	if err != nil {
		log.Fatal("ops #1 ", err)
	}

	//sort categories by last update
	sort.Slice(categories, func(i, j int) bool {
		return categories[i].LastUpdate < categories[j].LastUpdate
	})

	updateCategoryDataFile(categories, &appData)

	//sort categories by name
	sort.Slice(categories, func(i, j int) bool {
		return categories[i].Name < categories[j].Name
	})

	//update app.json
	appData.LastUpdate = getNowUnixTime()
	appDataJSON, _ := json.MarshalIndent(appData, "", "\t")
	saveFile("", appFileName, appDataJSON)

	//update readme
	readmeApp := getReadmeMD(categories)
	saveFile("", "README.md", readmeApp)
}

func getNowString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func getNowUnixTime() int64 {
	return time.Now().Unix()
}

func getReadmeMD(categories []Category) []byte {
	header := "# Go Libraries and Frameworks\n"
	header += "[![License](https://img.shields.io/github/license/IrvanAhmadP/Go-Libraries-and-Frameworks)](https://img.shields.io/github/license/IrvanAhmadP/Go-Libraries-and-Frameworks)\n\n"
	header += "List of Go frameworks, libraries and software inspired by [go-web-framework-stars](https://github.com/mingrammer/go-web-framework-stars).\n\n"
	header += "List of frameworks and libraries from [awesome-go](https://github.com/avelino/awesome-go).\n\n"
	readme := header

	readme += "### Contents\n"
	for _, category := range categories {
		categoryURL := strings.ReplaceAll(category.Name, " ", "-")
		categoryURL = "#" + strings.ToLower(categoryURL)
		readme += "* [" + category.Name + "](" + categoryURL + ")\n"
	}

	for _, category := range categories {
		categoryData := CategoryData{}
		categoryDataFileName := category.Name
		categoryDataFileName = strings.ReplaceAll(categoryDataFileName, " ", "-")
		categoryDataFileName = strings.ReplaceAll(categoryDataFileName, "(", "")
		categoryDataFileName = strings.ReplaceAll(categoryDataFileName, ")", "")
		categoryDataFileName = categoryDataFileName + ".json"
		categoryDataFile := readFile("data/" + categoryDataFileName)
		if categoryDataFile != nil {
			err := json.Unmarshal(categoryDataFile, &categoryData)
			if err != nil {
				fmt.Println(err)
			}

			categoryContent := getCategoryMD(category.Name, category.Description, categoryData.Libraries)
			readme += categoryContent + "\n\n"
		}
	}

	return []byte(readme)
}

func updateCategoryDataFile(categories []Category, appData *App) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("Recover")
		}
	}()

	for indexCategory, category := range categories {
		nowUnixTime := getNowUnixTime()
		repository := Library{}
		repositories := []Library{}

		listURL := ListURL{}
		listURLFile := readFile("url/" + category.Name + ".json")
		err := json.Unmarshal(listURLFile, &listURL)
		urls := listURL.URLs
		if err != nil {
			log.Fatal("ops #2 ", err)
		}

		if len(urls) <= appData.RateLimitRemaining || nowUnixTime > appData.RateLimitReset {
			fmt.Println("category: ", category.Name, " update")
			fmt.Println("libraries: ", len(urls))
			fmt.Println("RateLimitRemaining: ", appData.RateLimitRemaining)
			fmt.Println("nowUnixTime: ", nowUnixTime)
			fmt.Println("RateLimitReset: ", appData.RateLimitReset)

			categoryData := CategoryData{}
			categoryDataFileName := category.Name
			categoryDataFileName = strings.ReplaceAll(categoryDataFileName, " ", "-")
			categoryDataFileName = strings.ReplaceAll(categoryDataFileName, "(", "")
			categoryDataFileName = strings.ReplaceAll(categoryDataFileName, ")", "")
			categoryDataFileName = categoryDataFileName + ".json"
			categoryDataFile := readFile("data/" + categoryDataFileName)
			if categoryDataFile != nil {
				err = json.Unmarshal(categoryDataFile, &categoryData)
				if err != nil {
					fmt.Println(err)
				}
				repositories = categoryData.Libraries
			} else {
				categoryData.Name = category.Name
				categoryData.IndexLastCheck = 0
			}

			for indexURL, url := range urls {
				//pengecekan lanjut proses sebelumnya atau mulai dari awal
				if indexURL > categoryData.IndexLastCheck || categoryData.IndexLastCheck == 0 {
					categoryData.IndexLastCheck = indexURL
					if indexURL == (len(urls) - 1) {
						categoryData.IndexLastCheck = 0
					}

					resp, err := http.Get(url)
					if err != nil {
						panic("panic #1 : " + url + ":" + err.Error())
					}
					appData.RateLimit, _ = strconv.Atoi(resp.Header.Get("x-ratelimit-limit"))
					appData.RateLimitRemaining, _ = strconv.Atoi(resp.Header.Get("x-ratelimit-remaining"))
					appData.RateLimitReset, _ = strconv.ParseInt(resp.Header.Get("x-ratelimit-reset"), 10, 64)

					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						panic("panic #2 : " + err.Error())
					}

					if resp.StatusCode != 200 {
						panic("panic #3 : " + strconv.Itoa(resp.StatusCode) + " - " + url + " - " + string(body))
					}

					err = json.Unmarshal(body, &repository)
					if err != nil {
						panic("panic #4 : " + err.Error())
					}

					updateOldData := false
					for indexRepository, repositoryCheck := range repositories {
						if repositoryCheck.Id == repository.Id {
							repositories[indexRepository] = repository
							updateOldData = true
							break
						}
					}
					if !updateOldData {
						repositories = append(repositories, repository)
					}

					if appData.RateLimitRemaining == 0 {
						//save categoryData JSON #0
						//jika sudah limit
						categoryData.Libraries = repositories
						categoryData.LastUpdate = nowUnixTime
						saveCategoryDataJSON(categoryData, categoryDataFileName)
						break
					}
				}
			}
			//save categoryData JSON #1
			//jika belum limit
			categoryData.Libraries = repositories
			categoryData.LastUpdate = nowUnixTime
			saveCategoryDataJSON(categoryData, categoryDataFileName)

			categories[indexCategory].LastUpdate = nowUnixTime

		} else if len(urls) > 60 {
			fmt.Println("category: ", category.Name, " too many libraries ")
			fmt.Println("libraries: ", len(urls))
			fmt.Println("RateLimitRemaining: ", appData.RateLimitRemaining)
			fmt.Println("nowUnixTime: ", nowUnixTime)
			fmt.Println("RateLimitReset: ", appData.RateLimitReset)
		} else {
			fmt.Println("category: ", category.Name, " skip")
			fmt.Println("libraries: ", len(urls))
			fmt.Println("RateLimitRemaining: ", appData.RateLimitRemaining)
			fmt.Println("nowUnixTime: ", nowUnixTime)
			fmt.Println("RateLimitReset: ", appData.RateLimitReset)

		} //categories check end
		fmt.Println("========")
	}
}

func saveCategoryDataJSON(categoryData CategoryData, categoryDataFileName string) {
	libraries := categoryData.Libraries
	sort.Slice(libraries, func(i, j int) bool {
		return libraries[i].Stargazers > libraries[j].Stargazers
	})
	categoryDataJSON, _ := json.MarshalIndent(categoryData, "", "\t")
	saveFile("data", categoryDataFileName, categoryDataJSON)
}

func getCategoryMD(name string, description string, libraries []Library) string {
	headerTable := `## %s
%s

*Last Update: ` + getNowString() + `*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |`

	contentTable := "\n| [%s](%s) | %d | %d | %d | %s | %s | %s |"
	table := fmt.Sprintf(headerTable, name, description)
	for _, library := range libraries {
		repoURL := library.HomePageURL
		if library.HomePageURL == "" {
			repoURL = library.RepoURL
		}
		repoCreated := library.CreatedAt.Format("2006-01-02 15:04:05")
		repoUpdated := library.UpdatedAt.Format("2006-01-02 15:04:05")
		p := message.NewPrinter(language.English)
		table += p.Sprintf(contentTable, library.Name, repoURL, library.Stargazers, library.Forks, library.OpenIssues, library.Description, repoCreated, repoUpdated)
	}

	table += "\n</details>"
	return table
}

func readFile(path string) []byte {
	byteFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	return byteFile
}

func saveFile(path string, name string, content []byte) {
	fullPath := "./" + path + "/" + name
	if _, err := os.Stat(path); os.IsNotExist(err) && path != "" {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Fatal("ops #3 ", err.Error())
		}
	}

	// err := ioutil.WriteFile(fullPath, content, 0644)
	f, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal("ops #4 ", err.Error())
	}
	defer f.Close()

	f.Write(content)
}
