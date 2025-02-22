package nodes

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/tniedbala/tempe-go/tempe/api"
)

type textMarshaller struct {
	Text string `json:"text"`
}

type Text struct {
	text string
}

func NewText(text string) *Text {
	return &Text{text}
}

func (n *Text) Children() []api.TemplateNode {
	return []api.TemplateNode{}
}

func (n *Text) Render(opts api.Options, env api.Env, w io.StringWriter) error {
	_, err := w.WriteString(n.text)
	return err
}

func (n Text) Format() (string, string) {
	return "Text", fmt.Sprintf(`"%s"`, formatText(n.text))
}

func (n Text) String() string {
	return "Text{}"
}

func (n *Text) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonNodeSpec("Text", textMarshaller{
		Text: n.text,
	}))
}
