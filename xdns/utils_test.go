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
