package godoc

import (
	"fmt"
	"github.com/fabioxgn/go-bot"
	"github.com/fabioxgn/go-bot/web"
)

const (
	godocSiteURL    = "http://godoc.org"
	godocSearchURL  = "http://api.godoc.org/search?q=%s"
	noPackagesFound = "No packages found."
)

type godocResults struct {
	Results []struct {
		Path     string `json:"path"`
		Synopsis string `json:"synopsis"`
	} `json:"results"`
}

func searchGodoc(text string, get web.GetJSONFunc) (string, error) {
	if text == "" {
		return "", nil
	}

	data := &godocResults{}
	err := get(fmt.Sprintf(godocSearchURL, text), data)
	if err != nil {
		return "", err
	}

	if len(data.Results) == 0 {
		return noPackagesFound, nil
	}

	return fmt.Sprintf("%s %s/%s", data.Results[0].Synopsis, godocSiteURL, data.Results[0].Path), nil

}

func search(command *bot.Cmd) (msg string, err error) {
	return searchGodoc(command.FullArg, web.GetJSON)
}

func init() {
	bot.RegisterCommand(
		"godoc",
		"Searchs godoc.org and displays the first result.",
		"package name",
		search)
}
