package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIResponse DFC-ID用户手机号授权登录 API返回值
// alitrip.merchant.galaxy.wechat.user.authorize.login
//
// DFC-ID用户手机号授权登录
type AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIResponseModel is DFC-ID用户手机号授权登录 成功返回结果
type AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_wechat_user_authorize_login_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlitripMerchantGalaxyWechatUserAuthorizeLoginResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyWechatUserAuthorizeLoginAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIResponse)
	},
}

// GetAlitripMerchantGalaxyWechatUserAuthorizeLoginAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIResponse
func GetAlitripMerchantGalaxyWechatUserAuthorizeLoginAPIResponse() *AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIResponse {
	return poolAlitripMerchantGalaxyWechatUserAuthorizeLoginAPIResponse.Get().(*AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIResponse)
}

// ReleaseAlitripMerchantGalaxyWechatUserAuthorizeLoginAPIResponse 将 AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyWechatUserAuthorizeLoginAPIResponse(v *AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyWechatUserAuthorizeLoginAPIResponse.Put(v)
}
