package xdns

import (
	"testing"
	"time"
)

func TestDefaultGet(t *testing.T) {
	m := NewTTLMap(1 * time.Second)
	value := m.Get("test")
	if value != 0 {
		t.Errorf("expect %v, got %v", 0, value)
	}
}

func TestSetGet(t *testing.T) {
	m := NewTTLMap(1 * time.Second)
	m.Set("a", 1)
	value := m.Get("a")

	if value != 1 {
		t.Errorf("expect %v, got %v", 1, value)
	}
}

func TestTTL(t *testing.T) {
	m := NewTTLMap(1 * time.Microsecond)
	m.Set("a", 1)
	time.Sleep(2 * time.Millisecond)
	value := m.Get("a")

	if value != 0 {
		t.Errorf("expect %v, got %v", 0, value)
	}
}
