package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripCarOrderCompleteAPIRequest
服务完成API API请求
taobao.alitrip.car.order.complete

用来接收服务商订单流程完成信息 */
type TaobaoAlitripCarOrderCompleteAPIRequest struct {
	model.Params
	// 服务完成API
	_paramOrderComplete *OrderComplete
}

// New
