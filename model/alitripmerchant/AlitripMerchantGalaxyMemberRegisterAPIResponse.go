package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyMemberRegisterAPIResponse
星河-微信小程序会员注册 API返回值
alitrip.merchant.galaxy.member.register

星河产品=微信小程序注册雅高会员服务 */
type AlitripMerchantGalaxyMemberRegisterAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyMemberRegisterAPIResponseModel
}

// AlitripMerchantGalaxyMemberRegisterAPIResponseModel is 星河-微信小程序会员注册 成功返回结果
type AlitripMerchantGalaxyMemberRegisterAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_member_register_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyMemberRegisterResponse `json:"result,omitempty" xml:"result,omitempty"`
}
