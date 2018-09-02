package job

import "testing"

func TestStore(t *testing.T) {
	store := NewServicesStateStore()
	tables := []struct {
		name   string
		status bool
	}{
		{"disk", true},
		{"ram", false},
	}
	for _, table := range tables {
		store.SetState(table.name, table.status)
	}

	for _, table := range tables {
		result, err := store.GetState(table.name)
		if err != nil {
			t.Errorf(err.Error())
		}
		if result != table.status {
			t.Errorf("Output was incorrect, got %v want: %v", result, table.status)
		}
	}

	result, err := store.GetState("not in test")
	if result != false {
		t.Errorf("Output was incorrect, got %v want: %v", result, false)
	}
	if err == nil {
		t.Errorf("Output was incorrect, no error returned")
	}
}
