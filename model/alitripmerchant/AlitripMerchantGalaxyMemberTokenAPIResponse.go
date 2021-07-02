package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyMemberTokenAPIResponse 星河-校验token API返回值
// alitrip.merchant.galaxy.member.token
//
// 校验或者刷新token
type AlitripMerchantGalaxyMemberTokenAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyMemberTokenAPIResponseModel
}

// AlitripMerchantGalaxyMemberTokenAPIResponseModel is 星河-校验token 成功返回结果
type AlitripMerchantGalaxyMemberTokenAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_member_token_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyMemberTokenResponse `json:"result,omitempty" xml:"result,omitempty"`
}
