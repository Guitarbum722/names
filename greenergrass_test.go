package greenergrass

import "testing"

const targetTestVersion = 1

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}

func TestSeparateName(t *testing.T) {
	for _, ntc := range nameTestCases {
		got := SeparateName(ntc.input)
		if got.first != ntc.expected.first {
			t.Fatalf("First Name ::: SeparateName(%q) = %q, want %q", ntc.input, got.first, ntc.expected.first)
		}
		if got.middle != ntc.expected.middle {
			t.Fatalf("Middle Name ::: SeparateName(%q) = %q, want %q", ntc.input, got.middle, ntc.expected.middle)
		}
		if got.last != ntc.expected.last {
			t.Fatalf("Last Name ::: SeparateName(%q) = %q, want %q", ntc.input, got.last, ntc.expected.last)
		}
	}
}
