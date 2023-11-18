package marquee

import "testing"

func TestNew(t *testing.T) {
	m1 := New()
	m2 := New()
	if m1.id == m2.id {
		t.Errorf("Assigned IDs should be different, got %q and %q", m1.id, m2.id)
	}
}
