package alitripmerchant

import (
	"encoding/xml"

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

// AlitripMerchantGalaxyFavoriteQueryAPIResponseModel is 单酒店收藏状态查询 成功返回结果
type AlitripMerchantGalaxyFavoriteQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_favorite_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyFavoriteQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}
