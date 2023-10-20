package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxycouponinvalidlistAPIResponse 用户失效优惠券列表 API返回值
// alitrip.merchant.galaxy.coupon.invalid.list
//
// 雅高小程序用户失效优惠券列表
type AlitripmerchantgalaxycouponinvalidlistAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxycouponinvalidlistAPIResponseModel
}

// AlitripmerchantgalaxycouponinvalidlistAPIResponseModel is 用户失效优惠券列表 成功返回结果
type AlitripmerchantgalaxycouponinvalidlistAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_coupon_invalid_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripmerchantgalaxycouponinvalidlistResponse `json:"result,omitempty" xml:"result,omitempty"`
}
