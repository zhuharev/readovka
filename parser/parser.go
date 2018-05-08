// Copyright 2018 Kirill Zhuharev

package parser

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Parser struct {
	client *http.Client
}

type New struct {
	Title    string
	ID       int
	Body     string
	TopImage string
}

type News []New

func (p *Parser) DownloadNew(newID int) (New, error) {
	req, err := http.NewRequest("GET", "https://readovka.ru", nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36")

	resp, err := p.client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return nil, nil
}

func parseNew(id int, title string, body []byte) {

}
