package xxhash

import "testing"

func TestEscapes(t *testing.T) {
	cases := []struct {
		name string
		fn   func()
	}{
		{"Hash", func() {
			var buf [8]byte
			_ = Hash(buf[:])
		}},
		{"Hash128", func() {
			var buf [8]byte
			_ = Hash128(buf[:])
		}},
		{"HashString", func() {
			var buf [8]byte
			_ = HashString(string(buf[:]))
		}},
		{"HashString128", func() {
			var buf [8]byte
			_ = HashString128(string(buf[:]))
		}},
		{"HashSeed", func() {
			var buf [8]byte
			_ = HashSeed(buf[:], 1)
		}},
		{"Hash128Seed", func() {
			var buf [8]byte
			_ = Hash128Seed(buf[:], 1)
		}},
		{"HashStringSeed", func() {
			var buf [8]byte
			_ = HashStringSeed(string(buf[:]), 1)
		}},
		{"HashString128Seed", func() {
			var buf [8]byte
			_ = HashString128Seed(string(buf[:]), 1)
		}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if allocs := testing.AllocsPerRun(100, c.fn); allocs != 0 {
				t.Fatalf("allocations: %v, want 0", allocs)
			}
		})
	}
}
