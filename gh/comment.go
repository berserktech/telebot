package gh

import "fmt"

type Comment struct {
	Body    string
	HTMLURL string
}

// Message returns a string. // TODO finish this comment
func (c Comment) Format(kind string, s Sender) string {
	return fmt.Sprintf(`%s commented one %s with:

%s

%s`, s.Link(), kind, c.Body, c.HTMLURL)
}
