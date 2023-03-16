package jobs

import (
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/evanmcclur/mission-control/cron"
	"github.com/evanmcclur/mission-control/models"
)

// This cron is used to scrape featured articles off of space.com
type SpaceDotComJob struct {
	// When this cron job was last run
	LastRun time.Time
}

func (s *SpaceDotComJob) Run() (cron.Status, error) {
	res, err := http.Get("https://space.com")
	if err != nil {
		log.Print(err)
		// update last run timestamp
		s.LastRun = time.Now().UTC()
		return cron.Failed, err
	}
	defer res.Body.Close()
	// Check the response status code
	if res.StatusCode != 200 {
		// create http error to be returned
		err = NewHttpResponseError(res.StatusCode, res.Status)
		log.Print(err)
		// update last run timestamp
		s.LastRun = time.Now().UTC()
		return cron.Failed, err
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Print(err)
		// update last run timestamp
		s.LastRun = time.Now().UTC()
		return cron.Failed, err
	}

	articles := []*models.Article{}
	// Scrape articles from website and add to list
	doc.Find(".feature-block-item-wrapper > a.article-link").Each((func(i int, s *goquery.Selection) {
		article := models.NewArticle()
		// extract article link
		attr, exists := s.Attr("href")
		if !exists {
			// possibly return an error
			log.Print("Could not extract article link")
		}
		article.Link = attr

		// extract article title
		titleEl := s.Find(".article-name").First()
		if titleEl == nil {
			// possibly return an error
			log.Print("Could not extract article title")
		}
		article.Title = titleEl.Text()

		// extract article strapline
		straplineEl := s.Find(".article-strapline").First()
		if straplineEl == nil {
			// possibly return an error
			log.Print("Could not extract article strapline")
		}
		article.StrapLine = straplineEl.Text()
		articles = append(articles, article)
	}))

	// update last run timestamp
	s.LastRun = time.Now().UTC()
	return cron.Success, nil
}

func (s SpaceDotComJob) Key() cron.JobKey {
	return cron.SDC
}

func (s SpaceDotComJob) String() string {
	return "space.com"
}
