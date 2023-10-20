package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyMessageSubscriptionQueryAPIResponse 查询用户是否有模版ID权限 API返回值
// alitrip.merchant.galaxy.message.subscription.query
//
// 只是查询用户是否拥有权限ID
type AlitripMerchantGalaxyMessageSubscriptionQueryAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyMessageSubscriptionQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyMessageSubscriptionQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyMessageSubscriptionQueryAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyMessageSubscriptionQueryAPIResponseModel is 查询用户是否有模版ID权限 成功返回结果
type AlitripMerchantGalaxyMessageSubscriptionQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_message_subscription_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyMessageSubscriptionQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyMessageSubscriptionQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyMessageSubscriptionQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyMessageSubscriptionQueryAPIResponse)
	},
}

// GetAlitripMerchantGalaxyMessageSubscriptionQueryAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyMessageSubscriptionQueryAPIResponse
func GetAlitripMerchantGalaxyMessageSubscriptionQueryAPIResponse() *AlitripMerchantGalaxyMessageSubscriptionQueryAPIResponse {
	return poolAlitripMerchantGalaxyMessageSubscriptionQueryAPIResponse.Get().(*AlitripMerchantGalaxyMessageSubscriptionQueryAPIResponse)
}

// ReleaseAlitripMerchantGalaxyMessageSubscriptionQueryAPIResponse 将 AlitripMerchantGalaxyMessageSubscriptionQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyMessageSubscriptionQueryAPIResponse(v *AlitripMerchantGalaxyMessageSubscriptionQueryAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyMessageSubscriptionQueryAPIResponse.Put(v)
}
