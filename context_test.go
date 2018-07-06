package nova

import (
	"bytes"
	"errors"
	"net/http"
	"testing"
)

func TestContext_Next(t *testing.T) {
	nums := []byte{}
	var errOut error

	a := &Nova{
		Handlers: []HandlerFunc{
			func(c *Context) (err error) {
				c.Values["Key1"] = "AAA"
				nums = append(nums, 1)
				c.Next()
				nums = append(nums, 11)
				return
			},
			func(c *Context) (err error) {
				nums = append(nums, 2)
				c.Next()
				nums = append(nums, 12)
				return
			},
			func(c *Context) (err error) {
				nums = append(nums, 3)
				err = errors.New("ERROR" + c.Values["Key1"].(string))
				return
			},
		},
		ErrorHandler: func(c *Context, err error) {
			errOut = err
			return
		},
	}
	c := a.CreateContext(nil, &http.Request{})
	c.Next()
	if !bytes.Equal(nums, []byte{1, 2, 3, 12, 11}) || errOut.Error() != "ERRORAAA" {
		t.Error("bad Next()", nums)
	}
}

func TestContext_Panic(t *testing.T) {
	nums := []byte{}
	var errOut error

	a := &Nova{
		Handlers: []HandlerFunc{
			func(c *Context) (err error) {
				nums = append(nums, 1)
				c.Next()
				nums = append(nums, 11)
				return
			},
			func(c *Context) (err error) {
				nums = append(nums, 2)
				c.Next()
				nums = append(nums, 12)
				return
			},
			func(c *Context) (err error) {
				nums = append(nums, 3)
				panic("ERROR")
			},
		},
		ErrorHandler: func(c *Context, err error) {
			errOut = err
			return
		},
	}
	c := a.CreateContext(nil, &http.Request{})
	c.Next()
	if !bytes.Equal(nums, []byte{1, 2, 3, 12, 11}) || errOut.Error() != "panic: ERROR" {
		t.Error("bad Next()", nums)
	}
}
