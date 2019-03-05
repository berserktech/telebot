package gh

import "fmt"

type Content struct {
	Action  string
	Title   string
	HTMLURL string
	Body    string
}

// Format returns a string already formatted to be sent as a message.
func (c Content) Format(kind string, s Sender) string {
	var body string
	if c.Body != "" {
		body = fmt.Sprintf(" Details:\n%s", c.Body)
	}

	return fmt.Sprintf(
		"%s %s the %s: %s %s%s",
		s.Link(), c.Action, kind, c.Title, c.HTMLURL, body,
	)
}

// NotAllowed returns an error if the received Action is not valid.
func (c Content) NotAllowed() error {
	switch c.Action {
	case "labeled",
		"unlabeled",
		"assigned",
		"unassigned",
		"review_requested",
		"review_request_removed",
		"edited",
		"synchronize":
		return fmt.Errorf("gh: not allowed action, %s", c.Action)
	}

	return nil
}
