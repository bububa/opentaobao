package lstvending

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstvendingorderupdateAPIRequest 自动售货机订单物流信息回传 API请求
// alibaba.lst.vending.order.update
//
// 零售通与设备供应商进行订单对接，通过此接口回流订单物流信息。
type AlibabalstvendingorderupdateAPIRequest struct {
	model.Params
	// 零售通设备订单
	_vendingOrderDTO *VendingOrderDto
}

// NewAlibabalstvendingorderupdateRequest 初始化AlibabalstvendingorderupdateAPIRequest对象
func NewAlibabalstvendingorderupdateRequest() *AlibabalstvendingorderupdateAPIRequest {
	return &AlibabalstvendingorderupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstvendingorderupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.vending.order.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstvendingorderupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstvendingorderupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVendingOrderDTO is VendingOrderDTO Setter
// 零售通设备订单
func (r *AlibabalstvendingorderupdateAPIRequest) SetVendingOrderDTO(_vendingOrderDTO *VendingOrderDto) error {
	r._vendingOrderDTO = _vendingOrderDTO
	r.Set("vending_order_d_t_o", _vendingOrderDTO)
	return nil
}

// GetVendingOrderDTO VendingOrderDTO Getter
func (r AlibabalstvendingorderupdateAPIRequest) GetVendingOrderDTO() *VendingOrderDto {
	return r._vendingOrderDTO
}
