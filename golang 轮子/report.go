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

// Post POST爬取数据
func Post(url string,dT url.Values)string {
	ua := Readua()
	data := dT
	res, err := http.PostForm(url , data)
	res.Header.Set("User-Agent",ua)
	res.Header.Set("X-Forwarded-For","127.0.0.1")
	res.Header.Set("Referer","www.google.com")
	res.Header.Add("Authorization", "Bearer QSlHffXmCAILIOHNGXToq4LsP2yX64VQhEBZ7Ei4")
	if err != nil {
		fmt.Println("http post err",err)
		return "err"
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read err",err)
		return "err"
	}
	return string(b)
}
