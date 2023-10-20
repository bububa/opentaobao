package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyMemberLoginDerbyAPIResponse 小程序通过德比登入（会员认证） API返回值
// alitrip.merchant.galaxy.member.login.derby
//
// 会员认证(德比切换接口，包含授权接口)   返回认证及授权状态
type AlitripMerchantGalaxyMemberLoginDerbyAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyMemberLoginDerbyAPIResponseModel
}

// AlitripMerchantGalaxyMemberLoginDerbyAPIResponseModel is 小程序通过德比登入（会员认证） 成功返回结果
type AlitripMerchantGalaxyMemberLoginDerbyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_member_login_derby_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 携带优惠券试单结果
	Result *AlitripMerchantGalaxyMemberLoginDerbyResponse `json:"result,omitempty" xml:"result,omitempty"`
}
