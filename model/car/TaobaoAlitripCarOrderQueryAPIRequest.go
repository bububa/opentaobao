package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripCarOrderQueryAPIRequest
飞猪订单状态查询接口 API请求
taobao.alitrip.car.order.query

提供给直连商家查询在飞猪平台上产生的订单 */
type TaobaoAlitripCarOrderQueryAPIRequest struct {
	model.Params
	// 飞猪平台订单id
	_orderId string
}

// New
