package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripHotelDistributionOrderValidateAPIRequest
商旅酒店API分销下单前校验 API请求
alitrip.btrip.hotel.distribution.order.validate

商旅酒店API分销下单前校验 */
type AlitripBtripHotelDistributionOrderValidateAPIRequest struct {
	model.Params
	// 下单前校验入参
	_paramBtripHotelValidateOrderRq *BtripHotelValidateOrderRq
}

// New
