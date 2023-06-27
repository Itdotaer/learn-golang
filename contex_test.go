package main

import (
	"context"
	"testing"
	"time"
)

func TestContex(t *testing.T) {
	ctx := context.TODO()
	cancelCtx, cancelFunc := context.WithTimeout(ctx, time.Second*2)
	select {
	case <-cancelCtx.Done():
		println("超时", cancelCtx.Err().Error())
	}

	cancelFunc()
	println("取消ctx execute")

	ctx = context.WithValue(ctx, "demo", "djsfiojsdofj")
	println("ctx value", ctx.Value("demo").(string))
}
