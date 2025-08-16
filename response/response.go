package response

import (
	"encoding/json"

	"github.com/pratikiran/api-standard/context"
)

type Response struct {
	StatusCode int       `json:"-"`
	Reason     string    `json:"reason,omitempty"`
	Data       any       `json:"data,omitempty"`
	Warnings   []Warning `json:"warnings,omitempty"`
}

type Warning struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (resp *Response) EncodeToJSON(ctx *context.Context) []byte {
	bytes, err := json.Marshal(resp)
	if err != nil {
		ctx.Logger.Error("unable to marshal JSON: %s", err.Error())
	}

	return bytes
}

func (resp *Response) Error() string {
	return resp.Reason
}
