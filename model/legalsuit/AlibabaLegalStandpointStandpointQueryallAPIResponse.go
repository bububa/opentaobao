package legalsuit

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalstandpointstandpointqueryallAPIResponse 滑动查询口径 API返回值
// alibaba.legal.standpoint.standpoint.queryall
//
// 滑动查询口径
type AlibabalegalstandpointstandpointqueryallAPIResponse struct {
	model.CommonResponse
	AlibabalegalstandpointstandpointqueryallAPIResponseModel
}

// AlibabalegalstandpointstandpointqueryallAPIResponseModel is 滑动查询口径 成功返回结果
type AlibabalegalstandpointstandpointqueryallAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_standpoint_standpoint_queryall_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 口径
	Content []StandpointOutPutDto `json:"content,omitempty" xml:"content>standpoint_out_put_dto,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误编码
	ErrorCodeRes int64 `json:"error_code_res,omitempty" xml:"error_code_res,omitempty"`
	// 是否成功
	SuccessRes bool `json:"success_res,omitempty" xml:"success_res,omitempty"`
}
