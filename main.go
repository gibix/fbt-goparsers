package main

import (
	"log"
	"time"
	"github.com/gibix/fbtrex-cli/parser"
	"flag"
)

var parserType *string
var postType *string
var statusFlag *bool
var contentFlag *bool
var displayFlag *bool
var parserRun *bool
var commitFlag *bool

func init() {
	parserType = flag.String("parser", "postType", "parser name (postType, etc)")
	parserRun = flag.Bool("run", false, "run the parser handler")
	postType = flag.String("type", "feed", "post type (promoted, feed)")
	statusFlag = flag.Bool("status", false, "retrive the snippets content")
	contentFlag = flag.Bool("content", false, "retrive the snippets content")
	displayFlag = flag.Bool("display", false, "print snippet contents")
	commitFlag = flag.Bool("commit", false, "commit the parser's result")
}

func main() {
	flag.Parse()

	// TODO: something better?
	since := time.Date(2017, time.February, 28, 0, 0, 0, 0, time.UTC).Format("2006-01-02T15:04:05.511Z")
	until := time.Date(2017, time.February, 28, 10, 0, 0, 0, time.UTC).Format("2006-01-02T15:04:05.511Z")

	// define the query requirements
	p := parser.Parser {
		Profile: parser.SnippetInitProfile(*parserType, since, until, struct {
			Type string `json:"type"`
		}{
			Type: "feed",
		}),
		Snippets: []parser.SnippetContent{},
		Parsered:  []parser.SnippetResult{}}

	// get avaible snippets
	if (*statusFlag) {
		status := p.SnippetGetStatus()

		log.Println("Available:\t", status.Available)
		log.Println("Limit:\t", status.Limit)
	}

	// without parser key gets only 5 snippet
	if (*contentFlag) {
		_ = p.SnippetGetContent()
		if (*displayFlag) {
			log.Println(p.Snippets)
		}
	}

	// commits the parser's result
	if (*parserRun) {
		_ = p.ParserHandler()
		if (*commitFlag) {
			_ = p.CommitResult()
		}
	}
}
