package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxycouponvalidlistAPIResponse 用户有效优惠券列表 API返回值
// alitrip.merchant.galaxy.coupon.valid.list
//
// 雅高小程序用户有效优惠券列表
type AlitripmerchantgalaxycouponvalidlistAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxycouponvalidlistAPIResponseModel
}

// AlitripmerchantgalaxycouponvalidlistAPIResponseModel is 用户有效优惠券列表 成功返回结果
type AlitripmerchantgalaxycouponvalidlistAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_coupon_valid_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlitripmerchantgalaxycouponvalidlistResponse `json:"result,omitempty" xml:"result,omitempty"`
}
