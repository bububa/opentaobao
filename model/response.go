package model

import "encoding/xml"

type ErrorResponse struct {
	XMLName   xml.Name `xml:"error_response"`
	RequestId string   `json:"request_id,omitempty" xml:"request_id,omitempty"` // 平台颁发的每次请求访问的唯一标识
	Code      int      `json:"code,omitempty" xml:"code,omitempty"`             // 请求失败返回的错误码
	Msg       string   `json:"msg,omitempty" xml:"msg,omitempty"`               // 请求失败返回的错误信息
	SubCode   string   `json:"sub_code,omitempty" xml:"sub_code,omitempty"`     // 请求失败返回的子错误码
	SubMsg    string   `json:"sub_msg,omitempty" xml:"sub_msg,omitempty"`       // 请求失败返回的子错误信息
}

func (c ErrorResponse) Error() string {
	if c.SubMsg != "" {
		return c.SubMsg
	}
	return c.Msg
}

type CommonResponse struct {
	ErrorResponse *ErrorResponse `json:"error_response,omitempty" xml:"error_response,omitempty"` // 请求访问失败时返回的根节点
}

func (c CommonResponse) B043C16EB094F65A787F22E6AE0A10BCB7ABDE6D() {}

func (c CommonResponse) IsError() bool {
	return c.ErrorResponse != nil
}

func (c CommonResponse) Error() error {
	if c.ErrorResponse == nil {
		return nil
	}
	return c.ErrorResponse
}

type IResponse interface {
	B043C16EB094F65A787F22E6AE0A10BCB7ABDE6D()
	IsError() bool
	Error() error
}
