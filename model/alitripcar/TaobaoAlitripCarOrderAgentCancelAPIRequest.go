package alitripcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripCarOrderAgentCancelAPIRequest
司机或客服取消订单 API请求
taobao.alitrip.car.order.agent.cancel

司机或客服取消订单后通知飞猪订单取消 */
type TaobaoAlitripCarOrderAgentCancelAPIRequest struct {
	model.Params
	// 取消对象
	_paramOrderCancel *OrderCancel
}

// New
