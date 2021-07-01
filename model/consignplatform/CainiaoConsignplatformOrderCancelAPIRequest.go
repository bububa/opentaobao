package consignplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoConsignplatformOrderCancelAPIRequest
菜鸟发货工作台取消包裹以及订单 API请求
cainiao.consignplatform.order.cancel

菜鸟发货工作台，商家或者isv通过api取消包裹、回收单号，如果是裹裹运力会取消小件员上门。最后删除订单信息。 */
type CainiaoConsignplatformOrderCancelAPIRequest struct {
	model.Params
	// 取消参数
	_cancelRequest *OrderCancelRequest
}

// New
