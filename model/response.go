package model

type CommonResponse struct {
	RequestId     string `json:"request_id,omitempty"`     // 平台颁发的每次请求访问的唯一标识
	ErrorResponse string `json:"error_response,omitempty"` // 请求访问失败时返回的根节点
	Code          string `json:"code,omitempty"`           // 请求失败返回的错误码
	Msg           string `json:"msg,omitempty"`            // 请求失败返回的错误信息
	SubCode       string `json:"sub_code,omitempty"`       // 请求失败返回的子错误码
	SubMsg        string `json:"sub_msg,omitempty"`        // 请求失败返回的子错误信息
}

func (c CommonResponse) B043C16EB094F65A787F22E6AE0A10BCB7ABDE6D() {}

func (c CommonResponse) IsError() bool {
	return c.Code != "SUCCESS"
}

func (c CommonResponse) Error() string {
	return c.Msg
}

type IResponse interface {
	B043C16EB094F65A787F22E6AE0A10BCB7ABDE6D()
	IsError() bool
	Error() string
}