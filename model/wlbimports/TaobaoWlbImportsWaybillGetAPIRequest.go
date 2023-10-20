package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbimportswaybillgetAPIRequest 获取运单信息【云打印】 API请求
// taobao.wlb.imports.waybill.get
//
// 一般进口商家，获取订单的电子面单链接地址
type TaobaowlbimportswaybillgetAPIRequest struct {
	model.Params
	// 物流订单号
	_orderCode string
}

// NewTaobaowlbimportswaybillgetRequest 初始化TaobaowlbimportswaybillgetAPIRequest对象
func NewTaobaowlbimportswaybillgetRequest() *TaobaowlbimportswaybillgetAPIRequest {
	return &TaobaowlbimportswaybillgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbimportswaybillgetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.imports.waybill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbimportswaybillgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbimportswaybillgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCode is OrderCode Setter
// 物流订单号
func (r *TaobaowlbimportswaybillgetAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r TaobaowlbimportswaybillgetAPIRequest) GetOrderCode() string {
	return r._orderCode
}
