package main

import (
	"net/http/httptest"
	"os"
	"testing"
)

// To run this test, run the following command. Remember to replace the values with valid tokens and IDs:
// PORT=3000 TELEGRAM_TOKEN="TELEGRAM_TOKEN" TELEGRAM_CHAT_ID="TELEGRAM_CHAT_ID" INTEGRATION_TEST=true go test
func TestHandler(t *testing.T) {
	if os.Getenv("INTEGRATION_TEST") != "" {
		Handler(httptest.NewRecorder(), eventRequest("ping"))
	}
}
