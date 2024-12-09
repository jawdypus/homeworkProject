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
}

func Scrape() map[string]ClassDetails {

	// init homework map
	homework := make(map[string]ClassDetails)

	// get environmental variables
	loginUrl := os.Getenv("LOGIN_URL")
	dataUrl := os.Getenv("DATA_URL")
	username := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	// set variables
	var groups []string

	var groupName string
	var classTheme string
	var presence [][]string

	// set up chromedp
	browserPath := "C:\\Users\\raven\\AppData\\Local\\Thorium\\Application\\thorium.exe"

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.ExecPath(browserPath),
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
		tasks = getLastClassInfo(group, &groupName, &classTheme, &presence)

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
		}
	}

	return homework

}

func ScrapeF() map[string]ClassDetails {
	homework := make(map[string]ClassDetails)

	homework["Геймдизайн"] = ClassDetails{
		ClassTheme: "Відкритий урок",
		Presence: map[string]bool{
			"Юра":  true,
			"Іван": false,
		},
	}

	return homework
}
