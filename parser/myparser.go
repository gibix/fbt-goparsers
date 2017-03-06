package parser

type Result struct {
	Metadata string `json:"metadata"`
}

type Parser struct {
	ParserName string
	ParserKey  string
	Endpoint   string
	Profile    SnippetProfile
	Snippets   []SnippetContent
	Parsered   []SnippetResult
}

func (p *Parser) ParserHandler() error {
	for i, _ := range p.Snippets {
		p.Parsered[i] = SnippetResult{
			SnippetId:  p.Snippets[i].ObjectId,
			ParserName: p.ParserName,
			ParserKey:  p.ParserKey,
			Result: Result{
				Metadata: "meta"}}
	}

	return nil
}
