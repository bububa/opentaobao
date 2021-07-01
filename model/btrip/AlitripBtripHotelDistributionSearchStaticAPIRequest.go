package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripHotelDistributionSearchStaticAPIRequest
商旅酒店api分销-酒店静态信息接口 API请求
alitrip.btrip.hotel.distribution.search.static

商旅酒店api分销-酒店静态信息接口 */
type AlitripBtripHotelDistributionSearchStaticAPIRequest struct {
	model.Params
	// 基础信息入参
	_paramHotelInfoRQ *HotelInfoRq
}

// New
