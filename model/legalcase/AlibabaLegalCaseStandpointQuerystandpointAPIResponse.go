package legalcase

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalCaseStandpointQuerystandpointAPIResponse 主动查询口径 API返回值
// alibaba.legal.case.standpoint.querystandpoint
//
// 为法宝侧提供主动查询口径功能,有利于规范外部律师答辩口径.
type AlibabaLegalCaseStandpointQuerystandpointAPIResponse struct {
	model.CommonResponse
	AlibabaLegalCaseStandpointQuerystandpointAPIResponseModel
}

// AlibabaLegalCaseStandpointQuerystandpointAPIResponseModel is 主动查询口径 成功返回结果
type AlibabaLegalCaseStandpointQuerystandpointAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_case_standpoint_querystandpoint_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaLegalCaseStandpointQuerystandpointResult `json:"result,omitempty" xml:"result,omitempty"`
}
