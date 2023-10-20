package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmpointreversepointAPIResponse 积分消费回退 API返回值
// alibaba.alsc.crm.point.reversepoint
//
// 积分消费回退
type AlibabaalsccrmpointreversepointAPIResponse struct {
	model.CommonResponse
	AlibabaalsccrmpointreversepointAPIResponseModel
}

// AlibabaalsccrmpointreversepointAPIResponseModel is 积分消费回退 成功返回结果
type AlibabaalsccrmpointreversepointAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_point_reversepoint_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
