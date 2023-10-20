package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyShareGetAPIResponse 星河-获取小程序分享文案和图片 API返回值
// alitrip.merchant.galaxy.share.get
//
// 获取 雅高微信小程序分享素材文案和图片。
type AlitripMerchantGalaxyShareGetAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyShareGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyShareGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyShareGetAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyShareGetAPIResponseModel is 星河-获取小程序分享文案和图片 成功返回结果
type AlitripMerchantGalaxyShareGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_share_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyShareGetResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyShareGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyShareGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyShareGetAPIResponse)
	},
}

// GetAlitripMerchantGalaxyShareGetAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyShareGetAPIResponse
func GetAlitripMerchantGalaxyShareGetAPIResponse() *AlitripMerchantGalaxyShareGetAPIResponse {
	return poolAlitripMerchantGalaxyShareGetAPIResponse.Get().(*AlitripMerchantGalaxyShareGetAPIResponse)
}

// ReleaseAlitripMerchantGalaxyShareGetAPIResponse 将 AlitripMerchantGalaxyShareGetAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyShareGetAPIResponse(v *AlitripMerchantGalaxyShareGetAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyShareGetAPIResponse.Put(v)
}
