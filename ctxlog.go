package ctxlog

import (
	"context"
	"fmt"

	log "github.com/mycodesmells/golang-examples-log"
)

func Println(ctx context.Context, v ...interface{}) {
	label, ok := ctx.Value("label").(string)
	if ok && label != "" {
		v = append(v, 0)
		copy(v[1:], v[:])
		v[0] = fmt.Sprintf("[label=%s]", label)
	}
	log.Println(v...)
}
