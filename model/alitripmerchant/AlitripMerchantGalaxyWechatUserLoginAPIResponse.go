package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxywechatuserloginAPIResponse 微信小程序用户登录 API返回值
// alitrip.merchant.galaxy.wechat.user.login
//
// 微信小程序用户登录接口
type AlitripmerchantgalaxywechatuserloginAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxywechatuserloginAPIResponseModel
}

// AlitripmerchantgalaxywechatuserloginAPIResponseModel is 微信小程序用户登录 成功返回结果
type AlitripmerchantgalaxywechatuserloginAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_wechat_user_login_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripmerchantgalaxywechatuserloginResponse `json:"result,omitempty" xml:"result,omitempty"`
}
