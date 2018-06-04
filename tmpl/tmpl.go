package tmpl


type svg struct {
	Width string `xml:"width,attrs"`
	Height string `xml:"width,attrs"`
}

type Template struct {
	Height int
	Width int
	Elements []interface{}
	
}

func NewTemplate(path string) *Template {
	
	
	
}

func (t *Template) ReplaceText(id, text string) error {
	
}
