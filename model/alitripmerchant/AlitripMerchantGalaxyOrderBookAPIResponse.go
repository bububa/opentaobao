package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyOrderBookAPIResponse 星河-订单预订接口 API返回值
// alitrip.merchant.galaxy.order.book
//
// 为星河酒店解决方案的C端用户提供酒店预订能力
type AlitripMerchantGalaxyOrderBookAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyOrderBookAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyOrderBookAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyOrderBookAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyOrderBookAPIResponseModel is 星河-订单预订接口 成功返回结果
type AlitripMerchantGalaxyOrderBookAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_order_book_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyOrderBookResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyOrderBookAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyOrderBookAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyOrderBookAPIResponse)
	},
}

// GetAlitripMerchantGalaxyOrderBookAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyOrderBookAPIResponse
func GetAlitripMerchantGalaxyOrderBookAPIResponse() *AlitripMerchantGalaxyOrderBookAPIResponse {
	return poolAlitripMerchantGalaxyOrderBookAPIResponse.Get().(*AlitripMerchantGalaxyOrderBookAPIResponse)
}

// ReleaseAlitripMerchantGalaxyOrderBookAPIResponse 将 AlitripMerchantGalaxyOrderBookAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyOrderBookAPIResponse(v *AlitripMerchantGalaxyOrderBookAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyOrderBookAPIResponse.Put(v)
}
