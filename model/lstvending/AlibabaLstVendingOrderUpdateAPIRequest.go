package lstvending

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstVendingOrderUpdateAPIRequest) Reset() {
	r._vendingOrderDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstVendingOrderUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.vending.order.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstVendingOrderUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstVendingOrderUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaLstVendingOrderUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstVendingOrderUpdateRequest()
	},
}

// GetAlibabaLstVendingOrderUpdateRequest 从 sync.Pool 获取 AlibabaLstVendingOrderUpdateAPIRequest
func GetAlibabaLstVendingOrderUpdateAPIRequest() *AlibabaLstVendingOrderUpdateAPIRequest {
	return poolAlibabaLstVendingOrderUpdateAPIRequest.Get().(*AlibabaLstVendingOrderUpdateAPIRequest)
}

// ReleaseAlibabaLstVendingOrderUpdateAPIRequest 将 AlibabaLstVendingOrderUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstVendingOrderUpdateAPIRequest(v *AlibabaLstVendingOrderUpdateAPIRequest) {
	v.Reset()
	poolAlibabaLstVendingOrderUpdateAPIRequest.Put(v)
}
