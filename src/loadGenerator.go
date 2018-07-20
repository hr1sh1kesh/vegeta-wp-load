package src

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	vegeta "github.com/tsenart/vegeta/lib"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func init() {

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func GenerateLoadData(count int, duration int, api string, creds string) {
	auth := b64.StdEncoding.EncodeToString([]byte(creds))
	numberOfTargets := count * duration

	var targets []vegeta.Target

	rand.Seed(time.Now().UnixNano())
	ftype := []func(){generatePosts(api, auth, &targets), generateGETRequests(api, &targets)}

	for index := 0; index < numberOfTargets; index++ {
		generator := ftype[rand.Intn(len(ftype))]
		generator()
	}

	log.WithFields(log.Fields{"Number of targets generated": len(targets)}).Debug()

	rate := uint64(count) // per second
	du := time.Duration(duration) * time.Second
	attacker := vegeta.NewAttacker(vegeta.Workers(3), vegeta.Connections(1024))

	var metrics vegeta.Metrics
	var results vegeta.Results

	t := time.Now()
	loadTest := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	for res := range attacker.Attack(vegeta.NewStaticTargeter(targets...), rate, du, loadTest) {
		results.Add(res)
		metrics.Add(res)
	}
	var b bytes.Buffer
	plotReporter := vegeta.NewPlotReporter("wp-load-test"+loadTest, &results)
	err := plotReporter(&b)
	check(err)

	f, err := os.Create("load-test-output.html")
	check(err)
	defer f.Close()
	_, err = f.WriteString(b.String())
	check(err)

	metrics.Close()

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func generatePosts(api string, auth string, targets *[]vegeta.Target) func() {

	return func() {
		post := generateRandomPost()
		log.WithFields(log.Fields{"Request Body: ": post}).Debug()
		body, err := json.Marshal(post)
		check(err)
		var header = make(http.Header)
		header.Add("content-type", "application/json")
		header.Add("Authorization", "Basic "+auth)

		target := vegeta.Target{
			Method: http.MethodPost,
			URL:    api + "/wp-json/wp/v2/posts",
			Body:   body,
			Header: header,
		}
		addValue(targets, target)
	}
}

func generateGETRequests(api string, targets *[]vegeta.Target) func() {

	return func() {
		appURLs := []string{"/wp-json/wp/v2/posts", "/wp-json/wp/v2/users", "/wp-json/wp/v2/categories"}

		target := vegeta.Target{
			Method: http.MethodGet,
			URL:    api + appURLs[rand.Intn(len(appURLs))],
		}
		addValue(targets, target)
	}

}

func addValue(s *[]vegeta.Target, target vegeta.Target) {
	*s = append(*s, target)
	// fmt.Printf("In addValue: s is %v\n", s)
}

func generateRandomPost() Post {
	var post Post
	post.Title = RandStringBytes(25)
	post.Content = RandStringBytes(300)
	return post
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
