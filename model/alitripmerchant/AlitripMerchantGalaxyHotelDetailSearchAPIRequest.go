package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyHotelDetailSearchAPIRequest
星河-酒店详细信息搜索 API请求
alitrip.merchant.galaxy.hotel.detail.search

星河服务=获取雅高酒店详细信息 */
type AlitripMerchantGalaxyHotelDetailSearchAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 酒店详情入参
	_hotelDetailsParam *HotelDetailsParam
}

// New
