package filter

import "strings"

func In(field string, values ...string) string {
	if len(values) == 0 {
		return ""
	}

	// Each value should be quoted
	b := strings.Builder{}
	b.WriteString(field)
	b.WriteString(" IN (")
	for i, value := range values {
		b.WriteString("'")
		b.WriteString(value)
		b.WriteString("'")
		if i < len(values)-1 {
			b.WriteString(", ")
		}
	}
	b.WriteString(")")
	return b.String()
}
