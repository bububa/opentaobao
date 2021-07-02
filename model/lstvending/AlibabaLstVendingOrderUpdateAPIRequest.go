package lstvending

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstVendingOrderUpdateAPIRequest 自动售货机订单物流信息回传 API请求
// alibaba.lst.vending.order.update
//
// 零售通与设备供应商进行订单对接，通过此接口回流订单物流信息。
type AlibabaLstVendingOrderUpdateAPIRequest struct {
	model.Params
	// 零售通设备订单
	_vendingOrderDTO *VendingOrderDto
}

// NewAlibabaLstVendingOrderUpdateRequest 初始化AlibabaLstVendingOrderUpdateAPIRequest对象
func NewAlibabaLstVendingOrderUpdateRequest() *AlibabaLstVendingOrderUpdateAPIRequest {
	return &AlibabaLstVendingOrderUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstVendingOrderUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.vending.order.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstVendingOrderUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetVendingOrderDTO is VendingOrderDTO Setter
// 零售通设备订单
func (r *AlibabaLstVendingOrderUpdateAPIRequest) SetVendingOrderDTO(_vendingOrderDTO *VendingOrderDto) error {
	r._vendingOrderDTO = _vendingOrderDTO
	r.Set("vending_order_d_t_o", _vendingOrderDTO)
	return nil
}

// GetVendingOrderDTO VendingOrderDTO Getter
func (r AlibabaLstVendingOrderUpdateAPIRequest) GetVendingOrderDTO() *VendingOrderDto {
	return r._vendingOrderDTO
}
