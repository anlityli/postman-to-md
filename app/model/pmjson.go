package model

type Collection struct {
	Info *Info   `json:"info"`
	Item []*Item `json:"item"`
}

type Info struct {
	Name        string `json:"name"`
	Schema      string `json:"schema"`
	Version     string `json:"version"`
	Description string `json:"description"`
	PostmanID   string `json:"_postman_id"`
	ExporterId  string `json:"_exporter_id"`
}

type Item struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Variable    string      `json:"variable"`
	Request     *Request    `json:"request"`
	Response    []*Response `json:"response"`
	Item        []*Item     `json:"item"`
}

type Variable struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Value       string `json:"value"`
	Type        string `json:"type"` // enum: "string", "boolean", "any", "number"
	Name        string `json:"name"`
	Description string `json:"description"`
	System      bool   `json:"system"`
	Disabled    bool   `json:"disabled"`
}

type Request struct {
	Url         *Url      `json:"url"`
	Method      string    `json:"method"`
	Description string    `json:"description"`
	Header      []*Header `json:"header"`
	Body        *Body     `json:"body"`
}

type Response struct {
	Header []*Header `json:"header"`
	Body   string    `json:"body"`
}

type Header struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type Url struct {
	Raw   string    `json:"raw"`
	Query []*Header `json:"query"`
}

type Body struct {
	Mode string `json:"mode"` // "raw", "urlencoded", "formdata", "file", "graphql"
	Raw  string `json:"raw"`
}
