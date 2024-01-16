package scraper

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

const shareButtonSelector = "#tgl_basic_sh > div > ul > li.hasmore > div > ul > li:nth-child(3) > button"

func LoadTeamSeasonLog(team string, season int) (string, error) {
	url := fmt.Sprintf("https://www.basketball-reference.com/teams/%s/%d/gamelog", team, season)

	browser := rod.New().Timeout(time.Second * 3)
	err := browser.Connect()
	if err != nil {
		return "", fmt.Errorf("Failed to connect browser: %w", err)
	}
	defer browser.Close()

	page, err := browser.Page(proto.TargetCreateTarget{URL: url})
	if err != nil {
		return "", fmt.Errorf("Failed to open page: %w", err)
	}

	err = page.WaitIdle(time.Second)
	if err != nil {
		return "", fmt.Errorf("Error waiting for page idle: %w", err)
	}

	// Button is hidden on the page so custom JavaScript needs to be used to
	// access the element.
	//
	// In the JavaScript since the page takes a while to load, a Mutation
	// observer is used on the document to ensure that the JavaScript in the
	// page is loaded and the button exists on the page. Some sort of JavaScript
	// framework is being used which takes a while to load.
	_, err = page.Evaluate(&rod.EvalOptions{JS: fmt.Sprintf(`async () => {
		const selector = '%s';
		new Promise(resolve => {
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
	}`, shareButtonSelector)})
	if err != nil {
		return "", fmt.Errorf("error executing script: %w", err)
	}

	el, err := page.Element("#csv_tgl_basic")
	if err != nil {
		return "", fmt.Errorf("failed to find element: %w", err)
	}

	content, err := el.Text()
	if err != nil {
		return "", fmt.Errorf("error getting text: %w", err)
	}

	return content, nil
}
