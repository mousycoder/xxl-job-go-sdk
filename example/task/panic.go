package task

import (
	"context"
	xxl "github.com/mousycoder/xxl-job-go-sdk"
)

func Panic(cxt context.Context, param *xxl.RunReq) (msg string) {
	panic("test panic")
	return
}
