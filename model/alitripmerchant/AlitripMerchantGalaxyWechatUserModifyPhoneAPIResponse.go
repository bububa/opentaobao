package alitripmerchant

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlitripMerchantGalaxyWechatUserModifyPhoneAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyWechatUserModifyPhoneAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyWechatUserModifyPhoneAPIResponseModel is DFC-ID用户换绑手机号 成功返回结果
type AlitripMerchantGalaxyWechatUserModifyPhoneAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_wechat_user_modify_phone_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlitripMerchantGalaxyWechatUserModifyPhoneResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyWechatUserModifyPhoneAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyWechatUserModifyPhoneAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyWechatUserModifyPhoneAPIResponse)
	},
}

// GetAlitripMerchantGalaxyWechatUserModifyPhoneAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyWechatUserModifyPhoneAPIResponse
func GetAlitripMerchantGalaxyWechatUserModifyPhoneAPIResponse() *AlitripMerchantGalaxyWechatUserModifyPhoneAPIResponse {
	return poolAlitripMerchantGalaxyWechatUserModifyPhoneAPIResponse.Get().(*AlitripMerchantGalaxyWechatUserModifyPhoneAPIResponse)
}

// ReleaseAlitripMerchantGalaxyWechatUserModifyPhoneAPIResponse 将 AlitripMerchantGalaxyWechatUserModifyPhoneAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyWechatUserModifyPhoneAPIResponse(v *AlitripMerchantGalaxyWechatUserModifyPhoneAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyWechatUserModifyPhoneAPIResponse.Put(v)
}
