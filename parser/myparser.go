package parser

type MyNewContent struct {
    SnippetContent
}

type Parser struct {
    Profile SnippetProfile
    Snippets []SnippetContent
    Results []MyNewContent
}

func (p *Parser) ParserHandler() error {
	return nil
}
