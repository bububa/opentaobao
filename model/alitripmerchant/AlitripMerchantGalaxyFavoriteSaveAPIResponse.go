package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyFavoriteSaveAPIResponse 用户添加或移除收藏接口 API返回值
// alitrip.merchant.galaxy.favorite.save
//
// 用户可以点击收藏酒店，再次点击移除收藏的酒店
type AlitripMerchantGalaxyFavoriteSaveAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyFavoriteSaveAPIResponseModel
}

// AlitripMerchantGalaxyFavoriteSaveAPIResponseModel is 用户添加或移除收藏接口 成功返回结果
type AlitripMerchantGalaxyFavoriteSaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_favorite_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyFavoriteSaveResponse `json:"result,omitempty" xml:"result,omitempty"`
}
