package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyOrderCouponValidateAPIResponse 携带券的试单接口 API返回值
// alitrip.merchant.galaxy.order.coupon.validate
//
// 试单时可以使用优惠券，返回一个原价，一个折扣价
type AlitripMerchantGalaxyOrderCouponValidateAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyOrderCouponValidateAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyOrderCouponValidateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyOrderCouponValidateAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyOrderCouponValidateAPIResponseModel is 携带券的试单接口 成功返回结果
type AlitripMerchantGalaxyOrderCouponValidateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_order_coupon_validate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 携带优惠券试单结果
	Result *AlitripMerchantGalaxyOrderCouponValidateResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyOrderCouponValidateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyOrderCouponValidateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyOrderCouponValidateAPIResponse)
	},
}

// GetAlitripMerchantGalaxyOrderCouponValidateAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyOrderCouponValidateAPIResponse
func GetAlitripMerchantGalaxyOrderCouponValidateAPIResponse() *AlitripMerchantGalaxyOrderCouponValidateAPIResponse {
	return poolAlitripMerchantGalaxyOrderCouponValidateAPIResponse.Get().(*AlitripMerchantGalaxyOrderCouponValidateAPIResponse)
}

// ReleaseAlitripMerchantGalaxyOrderCouponValidateAPIResponse 将 AlitripMerchantGalaxyOrderCouponValidateAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyOrderCouponValidateAPIResponse(v *AlitripMerchantGalaxyOrderCouponValidateAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyOrderCouponValidateAPIResponse.Put(v)
}
