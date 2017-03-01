package parser

const ParserName = "MyParser"
const ParserKey = "myKey"

type Result struct {
	Metadata string `json:"metadata"`
}

type Parser struct {
    Profile SnippetProfile
    Snippets []SnippetContent
    Parsered []SnippetResult
}

func (p *Parser) ParserHandler() error {
	for i, _ := range p.Snippets {
		p.Parsered[i] = SnippetResult {
			SnippetId: p.Snippets[i].ObjectId,
			ParserName: ParserName,
			ParserKey: ParserKey,
			Result: Result {
				Metadata: "meta"}}
	}

	return nil
}
