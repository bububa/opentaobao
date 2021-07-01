package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripHotelDistributionSearchDetailAPIRequest
商旅酒店api分销-详情报价接口 API请求
alitrip.btrip.hotel.distribution.search.detail

商旅酒店api分销-详情报价接口 */
type AlitripBtripHotelDistributionSearchDetailAPIRequest struct {
	model.Params
	// 详情报价入参
	_paramHotelDetailRQ *HotelDetailRq
}

// New
