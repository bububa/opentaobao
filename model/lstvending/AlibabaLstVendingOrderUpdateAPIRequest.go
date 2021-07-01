package lstvending

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstVendingOrderUpdateAPIRequest
自动售货机订单物流信息回传 API请求
alibaba.lst.vending.order.update

零售通与设备供应商进行订单对接，通过此接口回流订单物流信息。 */
type AlibabaLstVendingOrderUpdateAPIRequest struct {
	model.Params
	// 零售通设备订单
	_vendingOrderDTO *VendingOrderDto
}

// New
