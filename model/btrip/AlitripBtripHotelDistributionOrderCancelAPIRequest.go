package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripHotelDistributionOrderCancelAPIRequest
商旅酒店API分销取消订单 API请求
alitrip.btrip.hotel.distribution.order.cancel

商旅酒店API分销取消订单 */
type AlitripBtripHotelDistributionOrderCancelAPIRequest struct {
	model.Params
	// 取消订单接口入参
	_paramBtripHotelOrderOperateRq *BtripHotelOrderOperateRq
}

// New
