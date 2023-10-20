package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyFavoriteQueryAPIResponse 单酒店收藏状态查询 API返回值
// alitrip.merchant.galaxy.favorite.query
//
// 返回用户对单个酒店的收藏状态
type AlitripMerchantGalaxyFavoriteQueryAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyFavoriteQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyFavoriteQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyFavoriteQueryAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyFavoriteQueryAPIResponseModel is 单酒店收藏状态查询 成功返回结果
type AlitripMerchantGalaxyFavoriteQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_favorite_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyFavoriteQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyFavoriteQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyFavoriteQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyFavoriteQueryAPIResponse)
	},
}

// GetAlitripMerchantGalaxyFavoriteQueryAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyFavoriteQueryAPIResponse
func GetAlitripMerchantGalaxyFavoriteQueryAPIResponse() *AlitripMerchantGalaxyFavoriteQueryAPIResponse {
	return poolAlitripMerchantGalaxyFavoriteQueryAPIResponse.Get().(*AlitripMerchantGalaxyFavoriteQueryAPIResponse)
}

// ReleaseAlitripMerchantGalaxyFavoriteQueryAPIResponse 将 AlitripMerchantGalaxyFavoriteQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyFavoriteQueryAPIResponse(v *AlitripMerchantGalaxyFavoriteQueryAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyFavoriteQueryAPIResponse.Put(v)
}
