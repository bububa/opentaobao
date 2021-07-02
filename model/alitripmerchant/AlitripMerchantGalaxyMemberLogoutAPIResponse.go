package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyMemberLogoutAPIResponse 星河-用户登出 API返回值
// alitrip.merchant.galaxy.member.logout
//
// 星河=微信小程序用户登出
type AlitripMerchantGalaxyMemberLogoutAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyMemberLogoutAPIResponseModel
}

// AlitripMerchantGalaxyMemberLogoutAPIResponseModel is 星河-用户登出 成功返回结果
type AlitripMerchantGalaxyMemberLogoutAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_member_logout_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyMemberLogoutResponse `json:"result,omitempty" xml:"result,omitempty"`
}
