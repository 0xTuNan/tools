package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Get 获取待爬取网站的html
func ValueGet(a string) string {
	ua := Readua()
	url := a
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", ua)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	Body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error", err)
	}
	doc := string(Body)
	return doc
}

// Post 发post包，三个参数，分别是url，body，cookie
func Post(a string, body string, cookie string) string {
	ua := Readua()
	url := a
	client := &http.Client{}

	payload := strings.NewReader(body)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("User-Agent", ua)
	req.Header.Add("Cookie", cookie)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	Body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error", err)
	}
	doc := string(Body)
	return doc
}
