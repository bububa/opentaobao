package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripHotelDistributionSearchLowPriceAPIRequest
商旅酒店api分销-酒店最低价 API请求
alitrip.btrip.hotel.distribution.search.low.price

商旅酒店api分销-酒店最低价 */
type AlitripBtripHotelDistributionSearchLowPriceAPIRequest struct {
	model.Params
	// 列表最低价入参
	_paramHotelSearchListRQ *HotelSearchListRq
}

// New
