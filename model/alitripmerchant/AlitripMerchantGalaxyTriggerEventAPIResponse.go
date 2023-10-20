package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyTriggerEventAPIResponse 抽奖活动自定义事件触发 API返回值
// alitrip.merchant.galaxy.trigger.event
//
// 自定义事件触发
type AlitripMerchantGalaxyTriggerEventAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyTriggerEventAPIResponseModel
}

// AlitripMerchantGalaxyTriggerEventAPIResponseModel is 抽奖活动自定义事件触发 成功返回结果
type AlitripMerchantGalaxyTriggerEventAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_trigger_event_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyTriggerEventResponse `json:"result,omitempty" xml:"result,omitempty"`
}
