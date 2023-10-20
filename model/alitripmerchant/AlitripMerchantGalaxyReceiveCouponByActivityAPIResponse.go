package alitripmerchant

import (
	"encoding/xml"

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

// AlitripMerchantGalaxyReceiveCouponByActivityAPIResponseModel is 按活动Id领取优惠券 成功返回结果
type AlitripMerchantGalaxyReceiveCouponByActivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_receive_coupon_by_activity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlitripMerchantGalaxyReceiveCouponByActivityResult `json:"result,omitempty" xml:"result,omitempty"`
}
