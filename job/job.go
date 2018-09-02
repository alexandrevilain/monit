package job

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type Job struct {
	Command string `yaml:"command"`
	Cron    string `yaml:"cron"`
	Test    string `yaml:"test"`
	Value   string `yaml:"value"`
}

func (j Job) ExecuteAndStore(serviceName string, store Store) func() {
	return func() {
		output, err := exec.Command("sh", "-c", j.Command).Output()
		if err != nil {
			store.SetState(j.Command, false)
		}
		result := j.outputToResult(strings.Trim(string(output), "\n"))
		store.SetState(serviceName, result)
	}
}

func (j Job) IsValid() error {
	// If it's not an equal or different test, si the value must be numeric
	if j.Test != "eq" && j.Test != "ne" {
		_, ok := getNumericValue(j.Value)
		if !ok {
			return fmt.Errorf("When using a test of type \"lt\" or \"gt\" the value must be of type number")
		}
	}
	return nil
}

func (j Job) outputToResult(output string) bool {
	isExpectedNumeric := true
	numericExpected, ok := getNumericValue(j.Value)
	if !ok {
		isExpectedNumeric = false
	}

	var numericOutput float64
	if isExpectedNumeric {
		numericOutput, ok = getNumericValue(output)
		if !ok {
			return false
		}
	}

	switch j.Test {
	case "eq", "=":
		return output == j.Value
	case "ne", "!=":
		return output != j.Value
	case "gt", ">":
		return numericOutput > numericExpected
	case "lt", "<":
		return numericOutput < numericExpected
	case "ge", ">=":
		return numericOutput >= numericExpected
	case "le", "<=":
		return numericOutput <= numericExpected
	default:
		return false
	}
}

func getNumericValue(s string) (float64, bool) {
	value, err := strconv.ParseFloat(s, 64)
	return value, err == nil
}
