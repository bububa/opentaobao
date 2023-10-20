package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyOrderCancelAPIResponse 星河-取消预订 API返回值
// alitrip.merchant.galaxy.order.cancel
//
// 雅高酒店用户使用该接口，取消酒店预订
type AlitripMerchantGalaxyOrderCancelAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyOrderCancelAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyOrderCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyOrderCancelAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyOrderCancelAPIResponseModel is 星河-取消预订 成功返回结果
type AlitripMerchantGalaxyOrderCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyOrderCancelResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyOrderCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyOrderCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyOrderCancelAPIResponse)
	},
}

// GetAlitripMerchantGalaxyOrderCancelAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyOrderCancelAPIResponse
func GetAlitripMerchantGalaxyOrderCancelAPIResponse() *AlitripMerchantGalaxyOrderCancelAPIResponse {
	return poolAlitripMerchantGalaxyOrderCancelAPIResponse.Get().(*AlitripMerchantGalaxyOrderCancelAPIResponse)
}

// ReleaseAlitripMerchantGalaxyOrderCancelAPIResponse 将 AlitripMerchantGalaxyOrderCancelAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyOrderCancelAPIResponse(v *AlitripMerchantGalaxyOrderCancelAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyOrderCancelAPIResponse.Put(v)
}
