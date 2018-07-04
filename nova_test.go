package nova

import (
	"bytes"
	"net/http"
	"net/url"
	"testing"
)

type responseWriter struct {
	buf    *bytes.Buffer
	status int
	header http.Header
}

func (r *responseWriter) Header() http.Header {
	return r.header
}

func (r *responseWriter) Write(p []byte) (int, error) {
	return r.buf.Write(p)
}

func (r *responseWriter) WriteHeader(statusCode int) {
	r.status = statusCode
}

func TestNova_ServeHTTP(t *testing.T) {
	r := &responseWriter{buf: &bytes.Buffer{}, header: http.Header{}}
	q := &http.Request{}
	q.URL, _ = url.Parse("/WORLD")
	a := New()
	a.Env = Production
	a.Use(func(ctx *Context) (err error) {
		ctx.Res.WriteHeader(200)
		ctx.Res.Write([]byte("HELLO" + ctx.Req.URL.Path + "/" + string(ctx.Env)))
		return
	})
	a.ServeHTTP(r, q)
	if r.buf.String() != "HELLO/WORLD/production" {
		t.Error("failed")
	}
}
