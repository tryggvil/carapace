package json

import (
	"encoding/json"
	"github.com/rsteube/carapace/internal/common"
)

type output struct {
	Nospace bool
	Values  []common.RawValue
}

func ActionRawValues(currentWord string, nospace bool, values common.RawValues) string {
	o := output{
		Nospace: nospace,
		Values:  values,
	}
	if m, err := json.Marshal(o); err == nil {
		return string(m)
	}
	return "{}"
}
