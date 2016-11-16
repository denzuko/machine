package versioncmp

import (
	"testing"
)

func TestCompareVersion(t *testing.T) {
	cases := []struct {
		v1, v2 string
		want   int
	}{
		{"1.12", "1.12", 0},
		{"1.0.0", "1", 0},
		{"1", "1.0.0", 0},
		{"1.05.00.0156", "1.0.221.9289", 1},
		{"1", "1.0.1", -1},
		{"1.0.1", "1", 1},
		{"1.0.1", "1.0.2", -1},
		{"1.0.2", "1.0.3", -1},
		{"1.0.3", "1.1", -1},
		{"1.1", "1.1.1", -1},
		{"1.a", "1.b", 0},
		{"1.a", "2.b", -1},
		{"1.1", "1.1.0", 0},
		{"1.1.1", "1.1.2", -1},
		{"1.1.2", "1.2", -1},
		{"1.12.1", "1.13.0-rc1", -1},
		{"1.13-rc1", "1.13.0-rc2", 0},
		{"1.13.0-rc1", "1.13.0-rc2", 0},
		{"1.13.0-rc1", "1.13.1-rc1", -1},
	}

	for _, tc := range cases {
		if got := compare(tc.v1, tc.v2); got != tc.want {
			t.Errorf("compare(%q, %q) == %d, want %d", tc.v1, tc.v2, got, tc.want)
		}
	}
}
