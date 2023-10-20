package alitripmerchant

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlitripMerchantGalaxyWechatLoginAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyWechatLoginAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyWechatLoginAPIResponseModel is 星河-用户使用微信登陆 成功返回结果
type AlitripMerchantGalaxyWechatLoginAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_wechat_login_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyWechatLoginResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyWechatLoginAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyWechatLoginAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyWechatLoginAPIResponse)
	},
}

// GetAlitripMerchantGalaxyWechatLoginAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyWechatLoginAPIResponse
func GetAlitripMerchantGalaxyWechatLoginAPIResponse() *AlitripMerchantGalaxyWechatLoginAPIResponse {
	return poolAlitripMerchantGalaxyWechatLoginAPIResponse.Get().(*AlitripMerchantGalaxyWechatLoginAPIResponse)
}

// ReleaseAlitripMerchantGalaxyWechatLoginAPIResponse 将 AlitripMerchantGalaxyWechatLoginAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyWechatLoginAPIResponse(v *AlitripMerchantGalaxyWechatLoginAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyWechatLoginAPIResponse.Put(v)
}
