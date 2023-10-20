package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyCardInfoAPIResponse 获取会员体系 API返回值
// alitrip.merchant.galaxy.card.info
//
// 星河=根据卡类型获取当前的会员体系
type AlitripMerchantGalaxyCardInfoAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyCardInfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyCardInfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyCardInfoAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyCardInfoAPIResponseModel is 获取会员体系 成功返回结果
type AlitripMerchantGalaxyCardInfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_card_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyCardInfoResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyCardInfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyCardInfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyCardInfoAPIResponse)
	},
}

// GetAlitripMerchantGalaxyCardInfoAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyCardInfoAPIResponse
func GetAlitripMerchantGalaxyCardInfoAPIResponse() *AlitripMerchantGalaxyCardInfoAPIResponse {
	return poolAlitripMerchantGalaxyCardInfoAPIResponse.Get().(*AlitripMerchantGalaxyCardInfoAPIResponse)
}

// ReleaseAlitripMerchantGalaxyCardInfoAPIResponse 将 AlitripMerchantGalaxyCardInfoAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyCardInfoAPIResponse(v *AlitripMerchantGalaxyCardInfoAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyCardInfoAPIResponse.Put(v)
}
