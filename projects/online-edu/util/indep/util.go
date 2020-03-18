package indep

import "encoding/json"

func Debug(src interface{}) string {
	buf, _ := json.MarshalIndent(src, "  ", "   ")
	return string(buf)
}
