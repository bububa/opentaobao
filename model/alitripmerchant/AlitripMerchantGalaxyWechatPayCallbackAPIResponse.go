package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyWechatPayCallbackAPIResponse 微信支付回调 API返回值
// alitrip.merchant.galaxy.wechat.pay.callback
//
// 微信支付回调
type AlitripMerchantGalaxyWechatPayCallbackAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyWechatPayCallbackAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyWechatPayCallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyWechatPayCallbackAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyWechatPayCallbackAPIResponseModel is 微信支付回调 成功返回结果
type AlitripMerchantGalaxyWechatPayCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_wechat_pay_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Result *AlitripMerchantGalaxyWechatPayCallbackResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyWechatPayCallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyWechatPayCallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyWechatPayCallbackAPIResponse)
	},
}

// GetAlitripMerchantGalaxyWechatPayCallbackAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyWechatPayCallbackAPIResponse
func GetAlitripMerchantGalaxyWechatPayCallbackAPIResponse() *AlitripMerchantGalaxyWechatPayCallbackAPIResponse {
	return poolAlitripMerchantGalaxyWechatPayCallbackAPIResponse.Get().(*AlitripMerchantGalaxyWechatPayCallbackAPIResponse)
}

// ReleaseAlitripMerchantGalaxyWechatPayCallbackAPIResponse 将 AlitripMerchantGalaxyWechatPayCallbackAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyWechatPayCallbackAPIResponse(v *AlitripMerchantGalaxyWechatPayCallbackAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyWechatPayCallbackAPIResponse.Put(v)
}
