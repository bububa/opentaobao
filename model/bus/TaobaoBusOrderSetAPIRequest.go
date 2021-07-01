package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusOrderSetAPIRequest
汽车票下单接口 API请求
taobao.bus.order.set

提供给汽车票商家进行下单 */
type TaobaoBusOrderSetAPIRequest struct {
	model.Params
	// 下单参数
	_paramB2BCreateOrderRQ *B2BCreateOrderRq
}

// New
