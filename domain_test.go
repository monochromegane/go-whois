package whois

import "testing"

func TestSld(t *testing.T) {
	d := domain{name: "hello.tokyo"}
	sld := d.sld()
	if sld != "hello" {
		t.Errorf("Expected string value 'hello', but %s", sld)
	}

	d.name = "hello.tokyo.jp"
	sld = d.sld()
	if sld != "hello" {
		t.Errorf("Expected string value 'hello', but %s", sld)
	}
}

func TestTld(t *testing.T) {
	d := domain{name: "hello.tokyo"}
	tld := d.tld()
	if tld != "tokyo" {
		t.Errorf("Expected string value 'tokyo', but %s", tld)
	}

	d.name = "hello.tokyo.jp"
	tld = d.tld()
	if tld != "tokyo.jp" {
		t.Errorf("Expected string value 'tokyo.jp', but %s", tld)
	}
}
