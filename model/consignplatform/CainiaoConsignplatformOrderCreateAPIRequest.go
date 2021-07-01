package consignplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoConsignplatformOrderCreateAPIRequest
菜鸟发货工作台创建订单 API请求
cainiao.consignplatform.order.create

菜鸟发货工作台，商家或者isv通过api进行订单写入操作 */
type CainiaoConsignplatformOrderCreateAPIRequest struct {
	model.Params
	// 订单创建入参
	_createRequest *OrderCreateRequest
}

// New
