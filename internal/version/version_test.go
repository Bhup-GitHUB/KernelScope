package version

import "testing"

func TestDefaultsAreNonEmpty(t *testing.T) {
	if Version == "" {
		t.Fatal("Version is empty")
	}

	if Commit == "" {
		t.Fatal("Commit is empty")
	}

	if Date == "" {
		t.Fatal("Date is empty")
	}
}
