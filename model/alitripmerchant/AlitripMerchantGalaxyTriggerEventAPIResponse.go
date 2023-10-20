package alitripmerchant

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlitripMerchantGalaxyTriggerEventAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyTriggerEventAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyTriggerEventAPIResponseModel is 抽奖活动自定义事件触发 成功返回结果
type AlitripMerchantGalaxyTriggerEventAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_trigger_event_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyTriggerEventResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyTriggerEventAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyTriggerEventAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyTriggerEventAPIResponse)
	},
}

// GetAlitripMerchantGalaxyTriggerEventAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyTriggerEventAPIResponse
func GetAlitripMerchantGalaxyTriggerEventAPIResponse() *AlitripMerchantGalaxyTriggerEventAPIResponse {
	return poolAlitripMerchantGalaxyTriggerEventAPIResponse.Get().(*AlitripMerchantGalaxyTriggerEventAPIResponse)
}

// ReleaseAlitripMerchantGalaxyTriggerEventAPIResponse 将 AlitripMerchantGalaxyTriggerEventAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyTriggerEventAPIResponse(v *AlitripMerchantGalaxyTriggerEventAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyTriggerEventAPIResponse.Put(v)
}
