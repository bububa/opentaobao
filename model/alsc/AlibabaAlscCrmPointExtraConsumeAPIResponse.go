package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmPointExtraConsumeAPIResponse 积分补扣 API返回值
// alibaba.alsc.crm.point.extra.consume
//
// 积分补扣
type AlibabaAlscCrmPointExtraConsumeAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmPointExtraConsumeAPIResponseModel
}

// AlibabaAlscCrmPointExtraConsumeAPIResponseModel is 积分补扣 成功返回结果
type AlibabaAlscCrmPointExtraConsumeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_point_extra_consume_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
