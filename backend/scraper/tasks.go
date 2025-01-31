package scraper

import (
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

func loginTask(loginUrl, username, password string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(loginUrl),

		chromedp.Sleep(2 * time.Second),

		chromedp.SendKeys(`#login`, username),
		chromedp.Click(`//*[@id="app"]/div/div/div[2]/form/div[2]/button`, chromedp.NodeVisible),

		chromedp.SendKeys(`#password`, password),

		chromedp.Click(`//*[@id="app"]/div/div/div[2]/form/div[3]/button`, chromedp.NodeVisible),

		chromedp.Sleep(3 * time.Second),
	}
}

func extractActiveGroups(dataUrl string, rows *[]string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(dataUrl),

		chromedp.WaitVisible(".table", chromedp.ByQuery),

		chromedp.Evaluate(`Array.from(document.querySelectorAll(".table tr"))
														.slice(1)
														.filter(row => row.children[10].innerText.includes("Група"))
														.map(row => row.children[0].innerText)`, &rows),
	}
}

func getLastClassInfo(groupId string, groupName *string, classTheme *string, presence *[][]string) chromedp.Tasks {
	baseGroupUrl := "https://lms.logikaschool.com/group/view/"
	lastClassPostfix := "#group-closest-lesson"
	groupUrl := fmt.Sprintf("%s%s%s", baseGroupUrl, groupId, lastClassPostfix)

	groupNameX := `/html/body/div[1]/div/div/div[3]/div/div/div[1]/div[1]/div/div[2]/div[1]/div`
	classThemeX := `/html/body/div[1]/div/div/div[3]/div/div/div[3]/div[4]/div/div[1]/div/h4/a`

	return chromedp.Tasks{
		chromedp.Navigate(groupUrl),

		chromedp.WaitVisible(groupNameX, chromedp.BySearch),

		chromedp.Text(groupNameX, groupName, chromedp.BySearch),

		chromedp.WaitVisible(".lessons-table", chromedp.ByQuery),

		chromedp.Text(classThemeX, classTheme, chromedp.BySearch),

		chromedp.Evaluate(`Array.from(document.querySelectorAll(".lessons-table tr td.lessons-table__meta-td div.lessons-table__td-inner")).map(div => {
  const name = div.querySelector("div#student-info div.lessons-table__meta-name a")?.innerText.trim();
  const presence = div.querySelector("div.lessons-table__meta-status div.student-presence div div button span")?.innerText.trim();
  return [name, presence]
}).filter(function( list ) {
  if (list[0] !== undefined)
   return list;
});`, &presence),
	}
}
