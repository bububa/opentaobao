package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripHotelDistributionOrderPayAPIRequest
商旅酒店分销订单支付 API请求
alitrip.btrip.hotel.distribution.order.pay

商旅酒店分销订单支付 */
type AlitripBtripHotelDistributionOrderPayAPIRequest struct {
	model.Params
	// 通知商旅支付成功接口参数
	_paramBtripHotelOrderOperateRq *BtripHotelOrderOperateRq
}

// New
