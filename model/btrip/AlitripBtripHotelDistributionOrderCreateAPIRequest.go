package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripHotelDistributionOrderCreateAPIRequest
商旅酒店分销-创建订单 API请求
alitrip.btrip.hotel.distribution.order.create

商旅酒店分销-创建订单 */
type AlitripBtripHotelDistributionOrderCreateAPIRequest struct {
	model.Params
	// 创建订单请求入参
	_paramBtripHotelCreateOrderRq *BtripHotelCreateOrderRq
}

// New
