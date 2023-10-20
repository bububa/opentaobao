package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyfavoritelistAPIResponse 用户收藏列表查询 API返回值
// alitrip.merchant.galaxy.favorite.list
//
// 让用户可以查询自己收藏的酒店列表
type AlitripmerchantgalaxyfavoritelistAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyfavoritelistAPIResponseModel
}

// AlitripmerchantgalaxyfavoritelistAPIResponseModel is 用户收藏列表查询 成功返回结果
type AlitripmerchantgalaxyfavoritelistAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_favorite_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripmerchantgalaxyfavoritelistResponse `json:"result,omitempty" xml:"result,omitempty"`
}
