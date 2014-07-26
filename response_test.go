package whois

import "testing"

func TestNewResponse(t *testing.T) {
	res := newResponse("com", "response")
	raw := res.Raw()
	if raw != "response" {
		t.Errorf("Expected response value 'response', but %s", raw)
	}
}
