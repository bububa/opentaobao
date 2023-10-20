package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyactivitycouponlistAPIResponse 用户领券中心列表 API返回值
// alitrip.merchant.galaxy.activity.coupon.list
//
// 雅高小程序用户领券中心列表
type AlitripmerchantgalaxyactivitycouponlistAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyactivitycouponlistAPIResponseModel
}

// AlitripmerchantgalaxyactivitycouponlistAPIResponseModel is 用户领券中心列表 成功返回结果
type AlitripmerchantgalaxyactivitycouponlistAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_activity_coupon_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripmerchantgalaxyactivitycouponlistResponse `json:"result,omitempty" xml:"result,omitempty"`
}
