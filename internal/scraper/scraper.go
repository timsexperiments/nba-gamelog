package scraper

import (
	"fmt"

	"github.com/go-rod/rod"
)

const shareButtonSelector = "#tgl_basic_sh > div > ul > li.hasmore > div > ul > li:nth-child(3) > button"

// LoadTeamSeasonLog loads the game log for a given team and season.
func LoadTeamSeasonLog(team string, season int) (string, error) {
	url := fmt.Sprintf("https://www.basketball-reference.com/teams/%s/%d/gamelog", team, season)

	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(url)
	page.MustWaitIdle()
	page.MustScreenshot("debug.png")

	// Button is hidden on the page so custom JavaScript needs to be used to
	// access the element.
	//
	// In the JavaScript since the page takes a while to load, a Mutation
	// observer is used on the document to ensure that the JavaScript in the
	// page is loaded and the button exists on the page. Some sort of JavaScript
	// framework is being used which takes a while to load.
	page.MustEval(fmt.Sprintf(`async () => {
		const selector = '%s';
		await new Promise(resolve => {
			const observer = new MutationObserver(mutations => {
				const el = document.querySelector(selector);
				if (el) {
					observer.disconnect();
					el.click();
					resolve();
				}
			});
	
			observer.observe(document.body, {
				childList: true,
				subtree: true
			});
		});
	}`, shareButtonSelector))

	content := page.MustElement("#csv_tgl_basic").MustText()

	return content, nil
}
