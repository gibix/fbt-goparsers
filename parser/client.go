package parser

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const ENDPOINT = "https://facebook.tracking.exposed"

// Simple parser constructor
func SnippetInitProfile(name string, since string, until string, require interface{}) SnippetProfile {
	return SnippetProfile{
		Requirements: require,
		Since:        since,
		Until:        until,
		ParserName:   name,
	}
}

// Returns number of available snippets for the query and limit
func (p *Parser) SnippetGetStatus() (int, int) {
	body, err := json.Marshal(p.Profile)
	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest("POST", ENDPOINT+"/api/v1/snippet/status", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}
	defer resp.Body.Close()

	var record SnippetStatus
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Fatal("Decode: ", err)
	}

	return record.Available, record.Limit
}

// Gets the available snippets in (*Parser).Snippets[]
func (p *Parser) SnippetGetContent() error {
	body, err := json.Marshal(p.Profile)
	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest("POST", ENDPOINT+"/api/v1/snippet/content", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&p.Snippets); err != nil {
		log.Fatal(err)
	}

	return err
}

// Commits the parser results from (*Parser).Parsered[]
func (p *Parser) CommitResult() error {
	body, err := json.Marshal(p.Parsered)
	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest("POST", ENDPOINT+"/api/v1/snippet/result", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&p.Snippets); err != nil {
		log.Fatal(err)
	}

	return err
}
