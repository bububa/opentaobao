package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyWechatUserLoginAPIResponse 微信小程序用户登录 API返回值
// alitrip.merchant.galaxy.wechat.user.login
//
// 微信小程序用户登录接口
type AlitripMerchantGalaxyWechatUserLoginAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyWechatUserLoginAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyWechatUserLoginAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyWechatUserLoginAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyWechatUserLoginAPIResponseModel is 微信小程序用户登录 成功返回结果
type AlitripMerchantGalaxyWechatUserLoginAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_wechat_user_login_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyWechatUserLoginResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyWechatUserLoginAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyWechatUserLoginAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyWechatUserLoginAPIResponse)
	},
}

// GetAlitripMerchantGalaxyWechatUserLoginAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyWechatUserLoginAPIResponse
func GetAlitripMerchantGalaxyWechatUserLoginAPIResponse() *AlitripMerchantGalaxyWechatUserLoginAPIResponse {
	return poolAlitripMerchantGalaxyWechatUserLoginAPIResponse.Get().(*AlitripMerchantGalaxyWechatUserLoginAPIResponse)
}

// ReleaseAlitripMerchantGalaxyWechatUserLoginAPIResponse 将 AlitripMerchantGalaxyWechatUserLoginAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyWechatUserLoginAPIResponse(v *AlitripMerchantGalaxyWechatUserLoginAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyWechatUserLoginAPIResponse.Put(v)
}
