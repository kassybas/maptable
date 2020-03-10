package vartable

import "fmt"

func unflattenPath(fields []interface{}, value interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	if len(fields) > 1 {
		res[fields[0].(string)] = unflattenNestedMap(fields[1:], value)
	} else {
		res[fields[0].(string)] = value
	}

	return res
}

func unflattenNestedMap(fields []interface{}, value interface{}) map[interface{}]interface{} {
	res := make(map[interface{}]interface{})
	if len(fields) == 1 {
		res[fields[0]] = value
		return res
	}
	res[fields[0]] = unflattenNestedMap(fields[1:], value)
	return res
}

func splitFields(path string) ([]string, error) {
	// errFmt := "could not parse variable path: %s\n%w"
	fields := []string{}
	buf := ""
	bracketsStarted := 0
	for _, c := range path {
		ch := string(c)
		switch ch {
		case ".":
			if bracketsStarted > 0 {
				// inside another bracket already
				buf += ch
			} else if len(buf) > 0 {
				fields = append(fields, buf)
				buf = ""
			}
		case "[":
			if bracketsStarted > 0 {
				// inside another bracket already
				buf += ch
			} else if len(buf) > 0 {
				fields = append(fields, buf)
				buf = ""
			}
			bracketsStarted++
		case "]":
			bracketsStarted--
			if bracketsStarted < 0 {
				return nil, fmt.Errorf("closing bracket without an opening: %s", path)
			}
			if bracketsStarted == 0 && len(buf) > 0 {
				fields = append(fields, buf)
				buf = ""
			} else {
				buf += ch
			}
		default:
			buf += ch
		}
	}
	if len(buf) != 0 {
		fields = append(fields, buf)
	}
	return fields, nil
}
