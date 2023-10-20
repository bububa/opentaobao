package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyOfferQueryAPIResponse 星河-offer查询 API返回值
// alitrip.merchant.galaxy.offer.query
//
// 根据offer的ID查询offer信息
type AlitripMerchantGalaxyOfferQueryAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyOfferQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyOfferQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyOfferQueryAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyOfferQueryAPIResponseModel is 星河-offer查询 成功返回结果
type AlitripMerchantGalaxyOfferQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_offer_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyOfferQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyOfferQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyOfferQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyOfferQueryAPIResponse)
	},
}

// GetAlitripMerchantGalaxyOfferQueryAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyOfferQueryAPIResponse
func GetAlitripMerchantGalaxyOfferQueryAPIResponse() *AlitripMerchantGalaxyOfferQueryAPIResponse {
	return poolAlitripMerchantGalaxyOfferQueryAPIResponse.Get().(*AlitripMerchantGalaxyOfferQueryAPIResponse)
}

// ReleaseAlitripMerchantGalaxyOfferQueryAPIResponse 将 AlitripMerchantGalaxyOfferQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyOfferQueryAPIResponse(v *AlitripMerchantGalaxyOfferQueryAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyOfferQueryAPIResponse.Put(v)
}
