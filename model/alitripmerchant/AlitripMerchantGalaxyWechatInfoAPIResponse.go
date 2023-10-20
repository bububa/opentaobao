package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyWechatInfoAPIResponse 星河-获取微信用户的信息 API返回值
// alitrip.merchant.galaxy.wechat.info
//
// 获取微信用户的openId和unionId
type AlitripMerchantGalaxyWechatInfoAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyWechatInfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyWechatInfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyWechatInfoAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyWechatInfoAPIResponseModel is 星河-获取微信用户的信息 成功返回结果
type AlitripMerchantGalaxyWechatInfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_wechat_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyWechatInfoResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyWechatInfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyWechatInfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyWechatInfoAPIResponse)
	},
}

// GetAlitripMerchantGalaxyWechatInfoAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyWechatInfoAPIResponse
func GetAlitripMerchantGalaxyWechatInfoAPIResponse() *AlitripMerchantGalaxyWechatInfoAPIResponse {
	return poolAlitripMerchantGalaxyWechatInfoAPIResponse.Get().(*AlitripMerchantGalaxyWechatInfoAPIResponse)
}

// ReleaseAlitripMerchantGalaxyWechatInfoAPIResponse 将 AlitripMerchantGalaxyWechatInfoAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyWechatInfoAPIResponse(v *AlitripMerchantGalaxyWechatInfoAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyWechatInfoAPIResponse.Put(v)
}
