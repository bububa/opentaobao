package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyMessageSubscriptionStorageAPIResponse 存储模版ID API返回值
// alitrip.merchant.galaxy.message.subscription.storage
//
// 消息订阅中的消息模版的存储
type AlitripMerchantGalaxyMessageSubscriptionStorageAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyMessageSubscriptionStorageAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyMessageSubscriptionStorageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyMessageSubscriptionStorageAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyMessageSubscriptionStorageAPIResponseModel is 存储模版ID 成功返回结果
type AlitripMerchantGalaxyMessageSubscriptionStorageAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_message_subscription_storage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyMessageSubscriptionStorageResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyMessageSubscriptionStorageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyMessageSubscriptionStorageAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyMessageSubscriptionStorageAPIResponse)
	},
}

// GetAlitripMerchantGalaxyMessageSubscriptionStorageAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyMessageSubscriptionStorageAPIResponse
func GetAlitripMerchantGalaxyMessageSubscriptionStorageAPIResponse() *AlitripMerchantGalaxyMessageSubscriptionStorageAPIResponse {
	return poolAlitripMerchantGalaxyMessageSubscriptionStorageAPIResponse.Get().(*AlitripMerchantGalaxyMessageSubscriptionStorageAPIResponse)
}

// ReleaseAlitripMerchantGalaxyMessageSubscriptionStorageAPIResponse 将 AlitripMerchantGalaxyMessageSubscriptionStorageAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyMessageSubscriptionStorageAPIResponse(v *AlitripMerchantGalaxyMessageSubscriptionStorageAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyMessageSubscriptionStorageAPIResponse.Put(v)
}
