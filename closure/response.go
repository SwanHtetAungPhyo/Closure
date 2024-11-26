package closure

import (
	"github.com/valyala/fasthttp"
	"github.com/goccy/go-json"
)

type JsonResponse struct {
    Status  int         `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}


func JSONfiy(ctx *fasthttp.RequestCtx, statusCode int, message string, data any) {
	response := JsonResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
	}

	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.SetStatusCode(statusCode)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		errorResponse := JsonResponse{
			Status:  fasthttp.StatusInternalServerError,
			Message: "Internal Server Error",
			Data:    nil,
		}
		errorJSON, _ := json.Marshal(errorResponse) 
		ctx.Response.SetBody(errorJSON)
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}


	ctx.Response.SetBody(jsonResponse)
}