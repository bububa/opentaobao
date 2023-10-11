package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyWechatUserModifyPhoneAPIResponse DFC-ID用户换绑手机号 API返回值
// alitrip.merchant.galaxy.wechat.user.modify.phone
//
// DFC-ID用户换绑手机号
type AlitripMerchantGalaxyWechatUserModifyPhoneAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyWechatUserModifyPhoneAPIResponseModel
}

// AlitripMerchantGalaxyWechatUserModifyPhoneAPIResponseModel is DFC-ID用户换绑手机号 成功返回结果
type AlitripMerchantGalaxyWechatUserModifyPhoneAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_wechat_user_modify_phone_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlitripMerchantGalaxyWechatUserModifyPhoneResponse `json:"result,omitempty" xml:"result,omitempty"`
}
