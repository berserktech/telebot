package gh

import "fmt"

type Status struct {
	State   string
	Message string
	HTMLURL string
}

// NotAllowed returns an error if the Status' State is not allowed to be
// handled.
func (s Status) NotAllowed() error {
	if s.State == "pending" {
		return fmt.Errorf("gh: not allowed status, pending")
	}

	return nil
}

// Format returns a string with a formatted message to be sent for this status
// with the passed sender.
func (status Status) Format(s Sender) string {
	return fmt.Sprintf(
		"`%s`: [%s](%s) by %s",
		status.State, status.Message, status.HTMLURL, s.Link(),
	)
}
