package utils

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
)

func FullPath(prefix , path string) string {
	return prefix + path
}

func Binder(ctx *fasthttp.RequestCtx,target any) error {
	body := ctx.PostBody()
	err := json.Unmarshal(body, target)
	if err 	!= nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return fmt.Errorf("error in json binding %s", err.Error())
	}
	return nil 
}