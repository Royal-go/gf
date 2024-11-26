package packed

import "github.com/gogf/gf/v2/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}