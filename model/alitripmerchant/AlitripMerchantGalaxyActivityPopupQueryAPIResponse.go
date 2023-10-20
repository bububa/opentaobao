package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyActivityPopupQueryAPIResponse 星河-获取雅高小程序营销抽奖首页弹窗 API返回值
// alitrip.merchant.galaxy.activity.popup.query
//
// 获取雅高微信小程序，营销抽奖首页弹窗数据。
type AlitripMerchantGalaxyActivityPopupQueryAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyActivityPopupQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyActivityPopupQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyActivityPopupQueryAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyActivityPopupQueryAPIResponseModel is 星河-获取雅高小程序营销抽奖首页弹窗 成功返回结果
type AlitripMerchantGalaxyActivityPopupQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_activity_popup_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyActivityPopupQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyActivityPopupQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyActivityPopupQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyActivityPopupQueryAPIResponse)
	},
}

// GetAlitripMerchantGalaxyActivityPopupQueryAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyActivityPopupQueryAPIResponse
func GetAlitripMerchantGalaxyActivityPopupQueryAPIResponse() *AlitripMerchantGalaxyActivityPopupQueryAPIResponse {
	return poolAlitripMerchantGalaxyActivityPopupQueryAPIResponse.Get().(*AlitripMerchantGalaxyActivityPopupQueryAPIResponse)
}

// ReleaseAlitripMerchantGalaxyActivityPopupQueryAPIResponse 将 AlitripMerchantGalaxyActivityPopupQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyActivityPopupQueryAPIResponse(v *AlitripMerchantGalaxyActivityPopupQueryAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyActivityPopupQueryAPIResponse.Put(v)
}
