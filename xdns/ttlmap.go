package xdns

import "time"

type Value struct {
	value int
	time  time.Time
}

type TTLMap struct {
	interval time.Duration
	data     map[string]Value
}

func NewTTLMap(interval time.Duration) TTLMap {
	return TTLMap{
		interval: interval,
	}
}

func (m *TTLMap) Get(key string) int {
	value, ok := m.data[key]

	if !ok {
		return 0
	}

	delta := time.Now().Sub(value.time)

	if delta > m.interval {
		return 0
	}

	return value.value
}

func (m *TTLMap) Set(key string, value int) {
	m.data[key] = Value{value: value, time: time.Now()}
}
