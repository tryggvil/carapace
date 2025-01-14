package carapace

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestActionImport(t *testing.T) {
	s := `
{
  "Version": "unknown",
  "Nospace": true,
  "RawValues": [
    {
      "Value": "positional1",
      "Display": "positional1",
      "Description": "",
      "Style": ""
    },
    {
      "Value": "p1",
      "Display": "p1",
      "Description": "",
      "Style": ""
    }
  ]
}`
	assertEqual(t, ActionValues("positional1", "p1").NoSpace().Invoke(Context{}), ActionImport([]byte(s)).Invoke(Context{}))
}

func TestActionFlags(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().BoolP("alpha", "a", false, "")
	cmd.Flags().BoolP("beta", "b", false, "")

	cmd.Flag("alpha").Changed = true
	a := actionFlags(cmd).Invoke(Context{CallbackValue: "-a"})
	assertEqual(t, ActionValuesDescribed("b", "").NoSpace().Invoke(Context{}).Prefix("-a"), a)
}

func TestActionExecCommandEnv(t *testing.T) {
	ActionExecCommand("env")(func(output []byte) Action {
		lines := strings.Split(string(output), "\n")
		for _, line := range lines {
			if strings.Contains(line, "carapace_TestActionExecCommand") {
				t.Error("should not contain env carapace_TestActionExecCommand")
			}
		}
		return ActionValues()
	}).Invoke(Context{})

	ActionExecCommand("env")(func(output []byte) Action {
		lines := strings.Split(string(output), "\n")
		for _, line := range lines {
			if line == "carapace_TestActionExecCommand=test" {
				return ActionValues()
			}
		}
		t.Error("should contain env carapace_TestActionExecCommand=test")
		return ActionValues()
	}).Invoke(Context{}.Setenv("carapace_TestActionExecCommand", "test"))
}
