package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyReceiveCouponByActivityAPIResponse 按活动Id领取优惠券 API返回值
// alitrip.merchant.galaxy.receive.coupon.by.activity
//
// 雅高小程序按活动Id领取优惠券
type AlitripMerchantGalaxyReceiveCouponByActivityAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyReceiveCouponByActivityAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyReceiveCouponByActivityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyReceiveCouponByActivityAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyReceiveCouponByActivityAPIResponseModel is 按活动Id领取优惠券 成功返回结果
type AlitripMerchantGalaxyReceiveCouponByActivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_receive_coupon_by_activity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlitripMerchantGalaxyReceiveCouponByActivityResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyReceiveCouponByActivityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyReceiveCouponByActivityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyReceiveCouponByActivityAPIResponse)
	},
}

// GetAlitripMerchantGalaxyReceiveCouponByActivityAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyReceiveCouponByActivityAPIResponse
func GetAlitripMerchantGalaxyReceiveCouponByActivityAPIResponse() *AlitripMerchantGalaxyReceiveCouponByActivityAPIResponse {
	return poolAlitripMerchantGalaxyReceiveCouponByActivityAPIResponse.Get().(*AlitripMerchantGalaxyReceiveCouponByActivityAPIResponse)
}

// ReleaseAlitripMerchantGalaxyReceiveCouponByActivityAPIResponse 将 AlitripMerchantGalaxyReceiveCouponByActivityAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyReceiveCouponByActivityAPIResponse(v *AlitripMerchantGalaxyReceiveCouponByActivityAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyReceiveCouponByActivityAPIResponse.Put(v)
}
