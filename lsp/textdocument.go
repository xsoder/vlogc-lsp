package lsp

type TextDocumentItem struct {
	URI        string `json:"uri"`
	LanguageId string `json:"languageid"`
	Version    int    `json:"version"`
	Text       string `json:"text"`
}

type TextDocumentIdentifier struct {
	URI string `json:"uri"`
}
type VersionTextDocumentIdentifier struct {
	TextDocumentIdentifier
	Version int `json:"version"`
}

type TextDocumentPositionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Position     Position               `json:"position"`
}
type Position struct {
	Line int `json:"line"`
	Character int `json:"charater"`
}

