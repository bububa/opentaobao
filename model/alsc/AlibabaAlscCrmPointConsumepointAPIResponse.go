package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmPointConsumepointAPIResponse 积分抵现 API返回值
// alibaba.alsc.crm.point.consumepoint
//
// 积分抵现
type AlibabaAlscCrmPointConsumepointAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmPointConsumepointAPIResponseModel
}

// AlibabaAlscCrmPointConsumepointAPIResponseModel is 积分抵现 成功返回结果
type AlibabaAlscCrmPointConsumepointAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_point_consumepoint_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
