package common

import (
	"fmt"
	"go-learn/projects/online-edu/util/indep"
)

func SaveError(e error, context ...map[string]interface{}) {
	if len(context) == 0 {
		context = []map[string]interface{}{
			map[string]interface{}{},
		}
	}
	context[0]["message"] = e.Error()

	fmt.Println(fmt.Sprintf("\n %s", indep.Debug(context)))
}
