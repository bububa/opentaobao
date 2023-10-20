package legalsuit

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalstandpointderivestandpointqueryAPIResponse 查询衍生口径 API返回值
// alibaba.legal.standpoint.derivestandpoint.query
//
// 查询衍生口径
type AlibabalegalstandpointderivestandpointqueryAPIResponse struct {
	model.CommonResponse
	AlibabalegalstandpointderivestandpointqueryAPIResponseModel
}

// AlibabalegalstandpointderivestandpointqueryAPIResponseModel is 查询衍生口径 成功返回结果
type AlibabalegalstandpointderivestandpointqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_standpoint_derivestandpoint_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCodeRes int64 `json:"error_code_res,omitempty" xml:"error_code_res,omitempty"`
	// 衍生口径
	Content *Page `json:"content,omitempty" xml:"content,omitempty"`
	// 成功
	SuccessRes bool `json:"success_res,omitempty" xml:"success_res,omitempty"`
}
