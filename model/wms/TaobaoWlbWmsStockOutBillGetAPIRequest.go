package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbwmsstockoutbillgetAPIRequest 通过订单号获取单个出库单发货信息 API请求
// taobao.wlb.wms.stock.out.bill.get
//
// 通过订单号获取单个出库单发货信息
type TaobaowlbwmsstockoutbillgetAPIRequest struct {
	model.Params
	// ERP订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
	_orderCode string
	// 菜鸟订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
	_cnOrderCode string
}

// NewTaobaowlbwmsstockoutbillgetRequest 初始化TaobaowlbwmsstockoutbillgetAPIRequest对象
func NewTaobaowlbwmsstockoutbillgetRequest() *TaobaowlbwmsstockoutbillgetAPIRequest {
	return &TaobaowlbwmsstockoutbillgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbwmsstockoutbillgetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wms.stock.out.bill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbwmsstockoutbillgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbwmsstockoutbillgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCode is OrderCode Setter
// ERP订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
func (r *TaobaowlbwmsstockoutbillgetAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r TaobaowlbwmsstockoutbillgetAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetCnOrderCode is CnOrderCode Setter
// 菜鸟订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
func (r *TaobaowlbwmsstockoutbillgetAPIRequest) SetCnOrderCode(_cnOrderCode string) error {
	r._cnOrderCode = _cnOrderCode
	r.Set("cn_order_code", _cnOrderCode)
	return nil
}

// GetCnOrderCode CnOrderCode Getter
func (r TaobaowlbwmsstockoutbillgetAPIRequest) GetCnOrderCode() string {
	return r._cnOrderCode
}
