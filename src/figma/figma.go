package figma

type FigmaNode struct {
	Id       string      `json:"id"`
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	Children []FigmaNode `json:"children"`
}

type FigmaDoc struct {
	Document FigmaNode `json:"document"`
}
