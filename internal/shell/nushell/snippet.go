// Package nushell provides Nushell completion
package nushell

import (
	"fmt"

	"github.com/rsteube/carapace/internal/uid"
	"github.com/spf13/cobra"
)

// Snippet creates the nushell completion script
func Snippet(cmd *cobra.Command) string {
	return fmt.Sprintf(`module completions {
    def "nu-complete %v" [line: string, pos: int] {
        $line | str substring ",$pos" | split row " " | %v _carapace nushell $in | from json
    }
    
    export extern "%v" [
        ...args: string@"nu-complete %v"
    ]
}
`, cmd.Name(), uid.Executable(), cmd.Name(), cmd.Name())
}
