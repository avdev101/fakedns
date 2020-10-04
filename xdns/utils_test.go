package xdns

import "testing"

func TestGetIps(t *testing.T) {
	ips := getIps("asdf-192.168.0.1-qqqaaazzz")

	if len(ips) != 1 {
		t.Errorf("len(ips) is %v", len(ips))
	}

	if ips[0] != "192.168.0.1" {
		t.Errorf("ip not match: %v", ips[0])
	}
}

func TestGetMultiIps(t *testing.T) {
	ips := getIps("asdf-11-192.168.0.1-192.168.1.1-12.-qqqaaazzz")

	if len(ips) != 2 {
		t.Errorf("len(ips) is %v", len(ips))
	}

	if ips[0] != "192.168.0.1" {
		t.Errorf("ip not match: %v", ips[0])
	}

	if ips[1] != "192.168.1.1" {
		t.Errorf("ip not match: %v", ips[1])
	}
}

func TestScheme(t *testing.T) {
	scheme := getScheme("1234-s011-1234a")
	expected := [3]int{0, 1, 1}

	if len(scheme) != len(expected) {
		t.Errorf("len mismatch")
	}

	if scheme[0] != expected[0] {
		t.Errorf("scheme[0] mismatch")
	}

	if scheme[1] != expected[1] {
		t.Errorf("scheme[1] mismatch")
	}

	if scheme[2] != expected[2] {
		t.Errorf("scheme[2] mismatch")
	}
}

func TestSchemeNoMatch(t *testing.T) {
	scheme := getScheme("1234s")
	if len(scheme) != 0 {
		t.Errorf("expected len %v got %v", 0, len(scheme))
	}
}
