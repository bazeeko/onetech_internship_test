package acmp_concurrent

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func diff(url string) float64 {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.AddCookie(&http.Cookie{Name: "English", Value: "1"})

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	r := regexp.MustCompile(`Difficulty: (\d+)%\)`)
	match := r.Find(data)

	var digits strings.Builder
	for _, v := range match {
		if v >= '0' && v <= '9' {
			digits.WriteByte(v)
		}
	}

	result, err := strconv.ParseFloat(digits.String(), 32)

	if result != 0 {
		return result
	}

	return -1
}

func Difficulties(urls []string) map[string]float64 {
	result := make(map[string]float64)
	mx := sync.Mutex{}
	wg := sync.WaitGroup{}

	for _, v := range urls {
		wg.Add(1)
		go func(v string) {
			mx.Lock()

			result[v] = diff(v)

			mx.Unlock()

			wg.Done()
		}(v)
	}
	wg.Wait()

	return result
}
