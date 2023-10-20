package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyWechatLoginAPIResponse 星河-用户使用微信登陆 API返回值
// alitrip.merchant.galaxy.wechat.login
//
// 星河产品=用户微信小程序登陆
type AlitripMerchantGalaxyWechatLoginAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyWechatLoginAPIResponseModel
}

// AlitripMerchantGalaxyWechatLoginAPIResponseModel is 星河-用户使用微信登陆 成功返回结果
type AlitripMerchantGalaxyWechatLoginAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_wechat_login_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyWechatLoginResponse `json:"result,omitempty" xml:"result,omitempty"`
}
