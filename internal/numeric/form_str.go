package numeric

import (
	"fmt"
	"strings"
)

// FormStr formats a numeric response into a slice of string parameters for the ircd.
func (r Response) FormStr(args ...interface{}) []string {
	res := fmt.Sprintf(formatStrings[r], args...)
	if res[0] == ':' {
		return []string{res[1:]}
	}

	split := strings.SplitN(res, " :", 2)
	params := strings.FieldsFunc(split[0], func(r rune) bool {
		return r == ' '
	})

	if len(split) == 1 {
		return params
	}

	return append(params, split[1])
}
