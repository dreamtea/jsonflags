package jsonflags

import (
	"flag"
	"testing"
)

func TestJsonFlags(t *testing.T) {
	args := []string{"-a", "a_value_from_arg", "-config=test_config.json"}
	flagSet := flag.NewFlagSet("jsonflags_test", flag.ContinueOnError)

	flagSet.String("config", "config.json", "config file path")

	flagSet.String("a", "a_default", "test flag A")
	flagSet.String("b", "b_default", "test flag B")
	flagSet.String("c", "c_default", "test flag C")
	flagSet.Int("d", 123, "test flag D")
	flagSet.Bool("e", true, "test flag E")
	flagSet.String("f", "", "test flag F")

	expected := map[string]interface{}{
		"a": "a_value_from_arg",
		"b": "b_value_from_json",
		"c": "c_default",
		"d": 456,
		"e": false,
		"f": "x,9.1,true",
	}

	err := ParseFlagSet(flagSet, args)
	if err != nil {
		t.Error(err)
		return
	}

	for name, value := range expected {
		f := flagSet.Lookup(name)
		if f == nil {
			t.Errorf("Missing flag %s", name)
			continue
		}
		actual := f.Value.(flag.Getter).Get()
		if actual != value {
			t.Fatalf("Incorrect value for %s: expected: %v actual: %v", name, value, actual)
		}
	}
}
