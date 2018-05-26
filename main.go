package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/json-iterator/go"
)

func main() {
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	var json GhostJSON
	err = jsoniter.Unmarshal(content, &json)
	if err != nil {
		panic(err)
	}
	data := json.Db[0].Data
	sourcePath := filepath.Join("./", "source")
	postsPath := filepath.Join(sourcePath, "_posts")
	draftPath := filepath.Join(sourcePath, "_drafts")
	os.MkdirAll(postsPath, os.ModePerm)
	os.MkdirAll(draftPath, os.ModePerm)
	for _, v := range data.Posts {
		var mobileDoc MobileDoc
		jsoniter.UnmarshalFromString(v.Mobiledoc, &mobileDoc)
		var card Card
		jsoniter.Unmarshal(mobileDoc.Cards[0][1], &card)

		publish, _ := time.Parse(time.RFC3339, v.PublishedAt)

		tags := data.getPostTags(v.ID)

		builder := bytes.Buffer{}
		builder.WriteString("---\n")
		builder.WriteString("title: \"" + v.Title + "\"\n")
		if v.PublishedAt != "" {
			builder.WriteString("date: " + publish.Local().Format("2006-01-02 15:04:05") + "\n")
		}
		// page don't has tags
		if len(tags) > 0 && v.Page != 1 {
			builder.WriteString("tags: \n")
			for _, v := range tags {
				builder.WriteString("- " + v + "\n")
			}
		}
		builder.WriteString("---\n")
		builder.WriteString(card.Markdown)
		fmt.Println(v.Title, v.Page, v.PublishedAt)
		if v.Status == "published" {
			if v.Page != 1 {
				ioutil.WriteFile(filepath.Join(postsPath, v.Slug+".md"), builder.Bytes(), os.ModePerm)
			} else {
				pagePath := filepath.Join(sourcePath, v.Slug)
				os.MkdirAll(pagePath, os.ModePerm)
				ioutil.WriteFile(filepath.Join(pagePath, "index.md"), builder.Bytes(), os.ModePerm)
			}
		}
		if v.Status == "draft" {
			ioutil.WriteFile(filepath.Join(draftPath, v.Slug+".md"), builder.Bytes(), os.ModePerm)
		}
	}
}
