package gh

import "fmt"

// Sender handles the author of the action.
type Sender struct {
	Login   string
	HTMLURL string
}

// Link returns a string with the URL for the Sender GitHub profile.
func (s Sender) Link() string {
	return fmt.Sprintf("[%s](%s)", s.Login, s.HTMLURL)
}
