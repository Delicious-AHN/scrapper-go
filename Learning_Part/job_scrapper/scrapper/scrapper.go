package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type extractedjob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

func Scrape(term string) {
	now := time.Now()
	baseURL := "https://kr.indeed.com/jobs?q=" + term + "&limit=50"
	var jobs []extractedjob
	c := make(chan []extractedjob)
	// totalpages := getPages()

	for i := 0; i < 9; i++ {
		go getPage(i, baseURL, c)
	}

	for i := 0; i < 9; i++ {
		extractedjobs := <-c
		jobs = append(jobs, extractedjobs...)
	}

	wErr := writeJobs(jobs)
	checkErr(wErr)

	fmt.Println("Done, Extracted", len(jobs))

	done := time.Now()

	fmt.Println(now.Sub(done))
}

func getPage(page int, url string, mainC chan<- []extractedjob) {
	var jobs []extractedjob
	c := make(chan extractedjob)
	pageURL := url + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting : ", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body) // Body is bite => memory
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body) // Body is bite => memory
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	return pages
}

func extractJob(card *goquery.Selection, c chan<- extractedjob) {
	id, _ := card.Attr("data-jk")
	title := CleanString(card.Find(".title>a").Text())
	location := CleanString(card.Find(".companyLocation").Text())
	salary := CleanString(card.Find(".salary-snippet").Text())
	summary := CleanString(card.Find(".summary").Text())
	c <- extractedjob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status Code : ", res.StatusCode)
	}
}

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func writeJobs(jobs []extractedjob) error {

	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "TITLE", "LOACTION", "SALARY", "SUMMARY"}

	wErr := w.Write(headers)
	checkErr(wErr)
	c := make(chan error)
	go csvWrite(jobs, w, c)

	for i := 0; i < len(jobs); i++ {
		<-c
	}
	return wErr
}

func csvWrite(jobs []extractedjob, w *csv.Writer, c chan<- error) {
	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
		c <- w.Write(jobSlice)
	}
}
