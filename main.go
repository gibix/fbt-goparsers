package main

import (
	"flag"
	"github.com/gibix/fbtrex-cli/parser"
	"github.com/spf13/viper"
	"log"
	"time"
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

const ISO8601 = "2006-01-02T15:04:05.511Z"

func main() {
	flag.Parse()

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("No configuration file")
	}

	// TODO: something better?
	since, err := time.Parse("2006-01-02 15:04:05", viper.GetString("since"))
	if err != nil {
		log.Fatal("Not possible to parse since date:", err)
	}
	until, err := time.Parse("2006-01-02 15:04:05", viper.GetString("until"))
	if err != nil {
		log.Fatal("Not possible to parse until date:", err)
	}

	// define the query requirements
	p := parser.Parser{
		Profile: parser.SnippetInitProfile(*parserType, since.String(), until.String(), struct {
			Type string `json:"type"`
		}{
			Type: "feed",
		}),
		Endpoint:   viper.GetString("endpoint"),
		Snippets:   []parser.SnippetContent{},
		Parsered:   []parser.SnippetResult{},
		ParserName: viper.GetString("parserKey"),
		ParserKey:  viper.GetString("parserName")}

	// get avaible snippets
	if *statusFlag {
		available, limit := p.SnippetGetStatus()

		log.Println("Available:\t", available)
		log.Println("Limit:\t", limit)
	}

	// without parser key gets only 5 snippet
	if *contentFlag {
		_ = p.SnippetGetContent()
		if *displayFlag {
			log.Println(p.Snippets)
		}
	}

	// commits the parser's result
	if *parserRun {
		_ = p.ParserHandler()
		if *commitFlag {
			_ = p.CommitResult()
		}
	}
}
