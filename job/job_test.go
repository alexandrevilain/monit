package job

import "testing"

func TestIsValue(t *testing.T) {
	tables := []struct {
		test  string
		value string
		err   bool
	}{
		{"eq", "test", false},
		{"ne", "test", false},
		{"gt", "80", false},
		{"gt", "test", true},
		{"le", "test", true},
	}

	for _, table := range tables {
		job := Job{Test: table.test, Value: table.value}
		result := job.IsValid()
		if result == nil && table.err {
			t.Errorf("Ouput was incorrect, got: %t", result)
		}
	}
}
func TestOutputToResult(t *testing.T) {
	tables := []struct {
		test     string
		expected string
		output   string
		result   bool
	}{
		{"eq", "test", "test", true},
		{"eq", "80", "80", true},
		{"eq", "80", "70", false},
		{"=", "80", "80", true},
		{"ne", "test", "no", true},
		{"ne", "test", "test", false},
		{"!=", "80", "70", true},
		{"gt", "60", "70", true},
		{"gt", "70", "60", false},
		{"gt", "test", "test", false},
		{">", "60", "70", true},
		{"lt", "70", "60", true},
		{"lt", "60", "70", false},
		{"<", "50", "60", false},
		{"ge", "70", "60", false},
		{"ge", "70", "70", true},
		{"le", "70", "60", true},
		{"le", "60", "70", false},
		{"le", "60", "60", true},
		{"unknown", "60", "60", false},
	}

	for _, table := range tables {
		job := Job{Test: table.test, Value: table.expected}
		result := job.outputToResult(table.output)
		if result != table.result {
			t.Errorf("Output was incorrect, got: %v, want: %v for test: %v", result, table.result, table.test)
		}
	}
}
