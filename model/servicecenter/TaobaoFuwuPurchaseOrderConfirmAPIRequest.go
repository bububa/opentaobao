package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofuwupurchaseorderconfirmAPIRequest 服务市场内购服务下单接口 API请求
// taobao.fuwu.purchase.order.confirm
//
// 通过传入服务市场商品的itemcode等信息，返回给服务商内购服务的下单链接
type TaobaofuwupurchaseorderconfirmAPIRequest struct {
	model.Params
	// 内购服务下单接口参数
	_paramOrderConfirmQueryDTO *OrderConfirmQueryDto
}

// NewTaobaofuwupurchaseorderconfirmRequest 初始化TaobaofuwupurchaseorderconfirmAPIRequest对象
func NewTaobaofuwupurchaseorderconfirmRequest() *TaobaofuwupurchaseorderconfirmAPIRequest {
	return &TaobaofuwupurchaseorderconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofuwupurchaseorderconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.fuwu.purchase.order.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofuwupurchaseorderconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofuwupurchaseorderconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOrderConfirmQueryDTO is ParamOrderConfirmQueryDTO Setter
// 内购服务下单接口参数
func (r *TaobaofuwupurchaseorderconfirmAPIRequest) SetParamOrderConfirmQueryDTO(_paramOrderConfirmQueryDTO *OrderConfirmQueryDto) error {
	r._paramOrderConfirmQueryDTO = _paramOrderConfirmQueryDTO
	r.Set("param_order_confirm_query_d_t_o", _paramOrderConfirmQueryDTO)
	return nil
}

// GetParamOrderConfirmQueryDTO ParamOrderConfirmQueryDTO Getter
func (r TaobaofuwupurchaseorderconfirmAPIRequest) GetParamOrderConfirmQueryDTO() *OrderConfirmQueryDto {
	return r._paramOrderConfirmQueryDTO
}
