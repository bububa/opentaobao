package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmPointExtrachargeAPIResponse
积分补录 API返回值
alibaba.alsc.crm.point.extracharge

积分补录 */
type AlibabaAlscCrmPointExtrachargeAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmPointExtrachargeAPIResponseModel
}

// AlibabaAlscCrmPointExtrachargeAPIResponseModel is 积分补录 成功返回结果
type AlibabaAlscCrmPointExtrachargeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_point_extracharge_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
