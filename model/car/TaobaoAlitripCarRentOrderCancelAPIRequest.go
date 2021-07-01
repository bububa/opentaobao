package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripCarRentOrderCancelAPIRequest
租车-取消订单 API请求
taobao.alitrip.car.rent.order.cancel

服务商主动取消用户订单或者拒绝取消订单. */
type TaobaoAlitripCarRentOrderCancelAPIRequest struct {
	model.Params
	// 取消请求对象
	_param0 *RentProviderCancelRequest
}

// New
