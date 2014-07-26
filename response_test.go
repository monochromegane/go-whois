package whois

import "testing"

func TestNewResponse(t *testing.T) {
	res := NewResponse("com", "response")
	raw := res.Raw()
	if raw != "response" {
		t.Errorf("Expected response value 'response', but %s", raw)
	}
}
