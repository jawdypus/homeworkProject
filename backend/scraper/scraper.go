package scraper

import (
	"context"
	"log"
	"os"

	"github.com/chromedp/chromedp"
)

type ClassDetails struct {
	ClassTheme string          `json:"class_theme"`
	Presence   map[string]bool `json:"presence"`
	Tasks      []string        `json:"tasks"`
}

func Scrape() map[string]ClassDetails {

	// init homework map
	homework := make(map[string]ClassDetails)

	// get environmental variables
	loginUrl := "https://lms.logikaschool.com/auth/v3/login"
	dataUrl := "https://lms.logikaschool.com/group?GroupSearch%5Bstatus%5D%5B0%5D=active&GroupSearch%5Bstatus%5D%5B1%5D=recruiting&presetType=all"
	username := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	// set variables
	var groups []string

	var groupName string
	var classTheme string
	var presence [][]string
	var classTasks []string

	// set up chromedp
	// browserPath := "C:\\Users\\raven\\AppData\\Local\\Thorium\\Application\\thorium.exe"

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-setuid-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("memory-pressure-off", true),
		//chromedp.ExecPath(browserPath),
	)

	allocatorCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocatorCtx)
	defer cancel()

	// define login and groups tasks
	tasks := chromedp.Tasks{}

	tasks = append(tasks, loginTask(loginUrl, username, password))
	tasks = append(tasks, extractActiveGroups(dataUrl, &groups))

	err := chromedp.Run(ctx, tasks)
	if err != nil {
		log.Fatalf("Failed to execute chromedp tasks: %v", err)
	}

	// get data from groups
	for _, group := range groups {

		groupName = ""
		classTheme = ""
		presence = [][]string{}
		classTasks = []string{}

		tasks = getLastClassInfo(group, &groupName, &classTheme, &presence)

		if err := chromedp.Run(ctx, tasks); err != nil {
			log.Fatalf("Failed to execute chromedp tasks: %v", err)
		}

		tasks = getLastClassTasks(group, &classTasks)
		if err := chromedp.Run(ctx, tasks); err != nil {
			log.Fatalf("Failed to execute chromedp tasks: %v", err)
		}

		presenceMap := make(map[string]bool)

		for _, student := range presence {
			if student[1] == "Був присутнім" {
				presenceMap[student[0]] = true
			} else if student[1] == "Відсутній" {
				presenceMap[student[0]] = false
			}
		}

		homework[groupName] = ClassDetails{
			ClassTheme: classTheme,
			Presence:   presenceMap,
			Tasks:      classTasks,
		}

	}

	return homework

}

func ScrapeF() map[string]ClassDetails {

	homework := make(map[string]ClassDetails)

	homework["Геймдизайн_Кременець_СБ_12:00"] = ClassDetails{
		ClassTheme: "Відкритий урок",
		Presence: map[string]bool{
			"Юра":  true,
			"Іван": false,
		},
		Tasks: []string{"Task1 Game ", "Task2 Game"},
	}

	homework["КГ_Кременець_НД_12.00"] = ClassDetails{
		ClassTheme: "Соц мережі",
		Presence: map[string]bool{
			"Міша": true,
			"Іра":  true,
		},
		Tasks: []string{"Task1 PC", "Task2 PC"},
	}

	return homework
}
