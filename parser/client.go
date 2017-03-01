package parser

import (
    "encoding/json"
    "log"
    "net/http"
    "bytes"
)

const ENDPOINT = "https://facebook.tracking.exposed"

func SnippetInitProfile (name string, since string, until string, require interface{}) SnippetProfile {
    return SnippetProfile {
        Requirements: require,
        Since : since,
        Until : until,
        ParserName:  name,
    }
}

func (p *Parser) SnippetGetStatus() SnippetStatus {
    body, err := json.Marshal(p.Profile)
    if err != nil {
        log.Println(err)
    }

    req, err := http.NewRequest("POST", ENDPOINT + "/api/v1/snippet/status", bytes.NewBuffer(body))
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

    return record
}

func (p *Parser) SnippetGetContent() error  {
    body, err := json.Marshal(p.Profile)
    if err != nil {
        log.Println(err)
    }

    req, err := http.NewRequest("POST", ENDPOINT + "/api/v1/snippet/content", bytes.NewBuffer(body))
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

func (è *Parser) SnippetCommit() error {
	return nil
}
