package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyFavoriteListAPIResponse 用户收藏列表查询 API返回值
// alitrip.merchant.galaxy.favorite.list
//
// 让用户可以查询自己收藏的酒店列表
type AlitripMerchantGalaxyFavoriteListAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyFavoriteListAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyFavoriteListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyFavoriteListAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyFavoriteListAPIResponseModel is 用户收藏列表查询 成功返回结果
type AlitripMerchantGalaxyFavoriteListAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_favorite_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyFavoriteListResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyFavoriteListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyFavoriteListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyFavoriteListAPIResponse)
	},
}

// GetAlitripMerchantGalaxyFavoriteListAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyFavoriteListAPIResponse
func GetAlitripMerchantGalaxyFavoriteListAPIResponse() *AlitripMerchantGalaxyFavoriteListAPIResponse {
	return poolAlitripMerchantGalaxyFavoriteListAPIResponse.Get().(*AlitripMerchantGalaxyFavoriteListAPIResponse)
}

// ReleaseAlitripMerchantGalaxyFavoriteListAPIResponse 将 AlitripMerchantGalaxyFavoriteListAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyFavoriteListAPIResponse(v *AlitripMerchantGalaxyFavoriteListAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyFavoriteListAPIResponse.Put(v)
}
