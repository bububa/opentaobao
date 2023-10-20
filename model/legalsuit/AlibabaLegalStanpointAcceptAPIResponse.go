package legalsuit

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalstanpointacceptAPIResponse 采纳口径 API返回值
// alibaba.legal.stanpoint.accept
//
// 采纳口径
type AlibabalegalstanpointacceptAPIResponse struct {
	model.CommonResponse
	AlibabalegalstanpointacceptAPIResponseModel
}

// AlibabalegalstanpointacceptAPIResponseModel is 采纳口径 成功返回结果
type AlibabalegalstanpointacceptAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_stanpoint_accept_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 状态code
	ErrorCodeRes int64 `json:"error_code_res,omitempty" xml:"error_code_res,omitempty"`
	// 是否成功
	SuccessRes bool `json:"success_res,omitempty" xml:"success_res,omitempty"`
	// 返回内容
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}
