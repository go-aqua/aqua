package nova

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"testing"
)

func TestContext_GetSetValue(t *testing.T) {
	req := &http.Request{}
	req = req.WithContext(context.WithValue(req.Context(), KeyType("a"), "b"))
	ctx := &Context{Req: req}
	if ctx.Value("a") != "b" {
		t.Error("failed 1")
	}
	ctx.Set("c", "d")
	if ctx.Req.Context().Value(KeyType("c")) != "d" {
		t.Error("failed 2")
	}
}

func TestContext_Next(t *testing.T) {
	nums := []byte{}
	var errOut error

	a := &Nova{
		Handlers: []HandlerFunc{
			func(ctx *Context) (err error) {
				nums = append(nums, 1)
				ctx.Next()
				nums = append(nums, 11)
				return
			},
			func(ctx *Context) (err error) {
				nums = append(nums, 2)
				ctx.Next()
				nums = append(nums, 12)
				return
			},
			func(ctx *Context) (err error) {
				nums = append(nums, 3)
				err = errors.New("ERROR")
				return
			},
		},
		ErrorHandler: func(ctx *Context, err error) {
			errOut = err
			return
		},
	}
	c := a.CreateContext(nil, &http.Request{})
	c.Next()
	if !bytes.Equal(nums, []byte{1, 2, 3, 12, 11}) || errOut.Error() != "ERROR" {
		t.Error("bad Next()", nums)
	}
}
