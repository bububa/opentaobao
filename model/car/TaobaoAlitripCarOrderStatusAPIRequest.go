package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripCarOrderStatusAPIRequest
商家订单状态改变通知接口（神州专车接口） API请求
taobao.alitrip.car.order.status

商家订单状态改变通知接口，神州专车专用接口！ */
type TaobaoAlitripCarOrderStatusAPIRequest struct {
	model.Params
	// 固定值：statusChanged
	_operation string
	// 飞猪订单ID
	_orderId string
	// 服务商ID
	_providerId string
	// 司机服务状态。arriving-司机已出发，arrived-司机已到达，serviceStarted-已开始服务，serviceFinished-已结束服务
	_status string
}

// New
