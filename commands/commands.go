package commands

import (
	"fmt"
	"net/url"
	"os"
	"time"
)

type Commands struct{}

func (c *Commands) Help() {
	// We want this method to show up in `help`/`list` results
	// But the actual logic for these commands is in `main.go`
}

func (c *Commands) List() {
	// We want this method to show up in `help`/`list` results
	// But the actual logic for these commands is in `main.go`
}

func (c *Commands) G(cmdArg string) string {
	return fmt.Sprintf("https://www.google.com/search?q=%s", url.QueryEscape(cmdArg))
}

func (c *Commands) Ags(cmdArg string) string {
	return fmt.Sprintf("https://github.com/search?q=repo%%3Aget-wrecked%%2Fmedal-electron+%s&type=code", url.QueryEscape(cmdArg))
}

func (c *Commands) Cal(cmdArg string) string {
	return "https://calendar.google.com/calendar/u/0/r"
}

func (c *Commands) Mail(cmdArg string) string {
	return "https://mail.google.com/mail/u/0/#inbox"
}

func (c *Commands) Chat(cmdArg string) string {
	return "https://chat.openai.com/"
}

func (c *Commands) Paste(cmdArg string) string {
	currentTime := time.Now()
	fileName := currentTime.Format("20060102150405") + ".txt"
	filePath := "./pastes/" + fileName

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		// TODO: write an error page
		return "https://www.youtube.com/watch?v=dQw4w9WgXcQ&ab_channel=RickAstley"
	}
	defer file.Close()

	_, err = file.WriteString(cmdArg)
	if err != nil {
		// TODO: write an error page
		return "https://www.youtube.com/watch?v=dQw4w9WgXcQ&ab_channel=RickAstley"
	}

	err = file.Sync()
	if err != nil {
		return "https://www.youtube.com/watch?v=dQw4w9WgXcQ&ab_channel=RickAstley"
	}

	return "https://www.youtube.com/watch?v=QTmkibc3WcI&ab_channel=Fai_rong"

	// TODO: figure out why tf this breaks??
	// return "file:///Users/braedon/Documents/gopherlol/pastes/" + fileName
}
