package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripCarOrderAcceptAPIRequest
确认接单 API请求
taobao.alitrip.car.order.accept

用来接收服务商确认接单信息 */
type TaobaoAlitripCarOrderAcceptAPIRequest struct {
	model.Params
	// 确认订单请求
	_paramOrderAccept *OrderAccept
}

// New
