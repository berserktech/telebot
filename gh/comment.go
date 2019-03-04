package gh

import "fmt"

type Comment struct {
	Body    string
	HTMLURL string
}

// Returns a formatted message saying who commented what, and where
func (c Comment) Format(kind string, s Sender) string {
	return fmt.Sprintf(`%s commented one %s with:

%s

%s`, s.Link(), kind, c.Body, c.HTMLURL)
}
