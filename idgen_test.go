package lya

import "testing"

func TestNew(t *testing.T) {
	u := NewId()
	// Check length
	if len(u) != StdLen {
		t.Fatalf("wrong length: expected %d, got %d", StdLen, len(u))
	}
	// Check that only allowed characters are present
	for _, c := range u {
		var present bool
		for _, a := range StdChars {
			if rune(a) == c {
				present = true
			}
		}
		if !present {
			t.Fatalf("chars not allowed in %q", u)
		}
	}
	// Generate 1000 uniuris and check that they are unique
	uris := make([]string, 1000)
	for i, _ := range uris {
		uris[i] = NewId()
		t.Log(uris[i])
	}
	for i, u := range uris {
		for j, u2 := range uris {
			if i != j && u == u2 {
				t.Fatalf("not unique: %d:%q and %d:%q", i, j, u, u2)
			}
		}
	}
}

func BenchmarkNewId(b *testing.B) {
	// Lets run the NewId() function a few times
	for i := 0; i < b.N; i++ {
		_ = NewId()
	}
}
