package parser

type SnippetProfile struct {
	Requirements interface{} `json:"requirements"`
	Since string `json:"since"`
	Until string `json:"until"`
	ParserName string `json:"parserName"`
}

type SnippetStatus struct {
	Available int `json:"available"`
	Limit int `json:"limit"`
}

type SnippetContent struct {
	ObjectId	string	`json:"_id"`
	SavingTime	string	`json:"savingTime"`
	Html		string	`json:"html"`
	Id		string	`json:"Id"`
	UserId		int	`json:"userId"`
	TimelineId	string	`json:"timelineId"`
	ImpressionId	string	`json:"impressionId"`
	PostType	bool	`json:"postType"`
	Type		string	`json"type"`
}

type SnippetContentPromotedTitle struct {
	SnippetContent
	PromotedTitle string `json:"promotedTitle"`
}

type SnippetContentPromotetLink struct {
	SnippetContent
	PromotedLink string `json:"promotedLink"`
}
