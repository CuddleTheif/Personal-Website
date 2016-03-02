package necrotheif

import (
	"net/http"
	"math/rand"
    "fmt"
	"strings"
	"appengine"
    	"appengine/urlfetch"
)

const characters = "abcdefghijklmnopqrstuvwxyz1234567890"

func code(w http.ResponseWriter, r *http.Request) {
	page := getCodePage()
	resp, _ := urlfetch.Client(appengine.NewContext(r)).Get(page)
    count := 0
	for strings.HasPrefix(resp.Header["Status"][0], "404") && count < 50{
	  page = getCodePage()
      fmt.Fprintln(w, page)
	  resp, _ = urlfetch.Client(appengine.NewContext(r)).Get(page)
      count += 1
	}
	http.Redirect(w, r, page, 301)
}

func getCodePage() string{

	page := "https://gist.github.com/anonymous/"
    entry := make([]byte, 20)
    for i := range entry {
        entry[i] = characters[rand.Intn(len(characters))]
    }
	return page+string(entry)
}