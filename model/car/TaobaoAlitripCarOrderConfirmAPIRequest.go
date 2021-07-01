package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripCarOrderConfirmAPIRequest
司机应答API API请求
taobao.alitrip.car.order.confirm

航旅事业群-度假事业部-旅行用车项目组对外部服务商提供的司机应答回调接口 */
type TaobaoAlitripCarOrderConfirmAPIRequest struct {
	model.Params
	// 请求对象
	_paramOrderConfirm *OrderConfirm
}

// New
