package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusOrderGetAPIRequest
汽车票订单查询 API请求
taobao.bus.order.get

商家汽车票订单查询 */
type TaobaoBusOrderGetAPIRequest struct {
	model.Params
	// 订单查询对象
	_paramB2BOrderQueryRQ *B2BOrderQueryRq
}

// New
