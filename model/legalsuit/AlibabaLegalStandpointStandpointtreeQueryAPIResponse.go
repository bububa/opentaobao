package legalsuit

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStandpointStandpointtreeQueryAPIResponse 查询口径树目录 API返回值
// alibaba.legal.standpoint.standpointtree.query
//
// 查询口径树目录
type AlibabaLegalStandpointStandpointtreeQueryAPIResponse struct {
	model.CommonResponse
	AlibabaLegalStandpointStandpointtreeQueryAPIResponseModel
}

// AlibabaLegalStandpointStandpointtreeQueryAPIResponseModel is 查询口径树目录 成功返回结果
type AlibabaLegalStandpointStandpointtreeQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_standpoint_standpointtree_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 口径树目录
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCodeRes int64 `json:"error_code_res,omitempty" xml:"error_code_res,omitempty"`
	// 是否成功
	SuccessRes bool `json:"success_res,omitempty" xml:"success_res,omitempty"`
}
