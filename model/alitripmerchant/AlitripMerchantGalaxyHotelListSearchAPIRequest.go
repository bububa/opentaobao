package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyHotelListSearchAPIRequest
星河-酒店列表页搜索 API请求
alitrip.merchant.galaxy.hotel.list.search

星河产品=酒店列表页搜索 */
type AlitripMerchantGalaxyHotelListSearchAPIRequest struct {
	model.Params
	// 商家租户id
	_tenantKey string
	// 请求参数
	_listSearchParam *ListSearchParam
}

// New
