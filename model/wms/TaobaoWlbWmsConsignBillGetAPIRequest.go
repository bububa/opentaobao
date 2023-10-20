package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbwmsconsignbillgetAPIRequest 获取销售订单发货信息 API请求
// taobao.wlb.wms.consign.bill.get
//
// 获取销售订单发货信息
type TaobaowlbwmsconsignbillgetAPIRequest struct {
	model.Params
	// 菜鸟订单编码,cnOrderCode与orderCode必须有一个值不可为空
	_cnOrderCode string
	// ERP订单编码,cnOrderCode与orderCode必须有一个值不可为空
	_orderCode string
}

// NewTaobaowlbwmsconsignbillgetRequest 初始化TaobaowlbwmsconsignbillgetAPIRequest对象
func NewTaobaowlbwmsconsignbillgetRequest() *TaobaowlbwmsconsignbillgetAPIRequest {
	return &TaobaowlbwmsconsignbillgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbwmsconsignbillgetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wms.consign.bill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbwmsconsignbillgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbwmsconsignbillgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCnOrderCode is CnOrderCode Setter
// 菜鸟订单编码,cnOrderCode与orderCode必须有一个值不可为空
func (r *TaobaowlbwmsconsignbillgetAPIRequest) SetCnOrderCode(_cnOrderCode string) error {
	r._cnOrderCode = _cnOrderCode
	r.Set("cn_order_code", _cnOrderCode)
	return nil
}

// GetCnOrderCode CnOrderCode Getter
func (r TaobaowlbwmsconsignbillgetAPIRequest) GetCnOrderCode() string {
	return r._cnOrderCode
}

// SetOrderCode is OrderCode Setter
// ERP订单编码,cnOrderCode与orderCode必须有一个值不可为空
func (r *TaobaowlbwmsconsignbillgetAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r TaobaowlbwmsconsignbillgetAPIRequest) GetOrderCode() string {
	return r._orderCode
}
