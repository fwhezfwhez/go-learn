package indep

import (
	"encoding/json"
	"fmt"
)

type InnerContext struct {
	RequestBody   json.RawMessage
	JSONDoNothing bool
}

func DefaultInnerContext(param map[string]interface{}) *InnerContext {
	b, _ := json.Marshal(param)
	return &InnerContext{
		RequestBody: json.RawMessage(b),
	}
}
func (c *InnerContext) Bind(dest interface{}) error {
	return json.Unmarshal(c.RequestBody, dest)
}

func (c *InnerContext) JSON(status int, resp interface{}) {
	if c.JSONDoNothing == false {
		fmt.Println(Debug(resp))
	}
}
