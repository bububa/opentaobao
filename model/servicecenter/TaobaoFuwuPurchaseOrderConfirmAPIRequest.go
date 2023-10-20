package servicecenter

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFuwuPurchaseOrderConfirmAPIRequest 服务市场内购服务下单接口 API请求
// taobao.fuwu.purchase.order.confirm
//
// 通过传入服务市场商品的itemcode等信息，返回给服务商内购服务的下单链接
type TaobaoFuwuPurchaseOrderConfirmAPIRequest struct {
	model.Params
	// 内购服务下单接口参数
	_paramOrderConfirmQueryDTO *OrderConfirmQueryDto
}

// NewTaobaoFuwuPurchaseOrderConfirmRequest 初始化TaobaoFuwuPurchaseOrderConfirmAPIRequest对象
func NewTaobaoFuwuPurchaseOrderConfirmRequest() *TaobaoFuwuPurchaseOrderConfirmAPIRequest {
	return &TaobaoFuwuPurchaseOrderConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFuwuPurchaseOrderConfirmAPIRequest) Reset() {
	r._paramOrderConfirmQueryDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFuwuPurchaseOrderConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.fuwu.purchase.order.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFuwuPurchaseOrderConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFuwuPurchaseOrderConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOrderConfirmQueryDTO is ParamOrderConfirmQueryDTO Setter
// 内购服务下单接口参数
func (r *TaobaoFuwuPurchaseOrderConfirmAPIRequest) SetParamOrderConfirmQueryDTO(_paramOrderConfirmQueryDTO *OrderConfirmQueryDto) error {
	r._paramOrderConfirmQueryDTO = _paramOrderConfirmQueryDTO
	r.Set("param_order_confirm_query_d_t_o", _paramOrderConfirmQueryDTO)
	return nil
}

// GetParamOrderConfirmQueryDTO ParamOrderConfirmQueryDTO Getter
func (r TaobaoFuwuPurchaseOrderConfirmAPIRequest) GetParamOrderConfirmQueryDTO() *OrderConfirmQueryDto {
	return r._paramOrderConfirmQueryDTO
}

var poolTaobaoFuwuPurchaseOrderConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFuwuPurchaseOrderConfirmRequest()
	},
}

// GetTaobaoFuwuPurchaseOrderConfirmRequest 从 sync.Pool 获取 TaobaoFuwuPurchaseOrderConfirmAPIRequest
func GetTaobaoFuwuPurchaseOrderConfirmAPIRequest() *TaobaoFuwuPurchaseOrderConfirmAPIRequest {
	return poolTaobaoFuwuPurchaseOrderConfirmAPIRequest.Get().(*TaobaoFuwuPurchaseOrderConfirmAPIRequest)
}

// ReleaseTaobaoFuwuPurchaseOrderConfirmAPIRequest 将 TaobaoFuwuPurchaseOrderConfirmAPIRequest 放入 sync.Pool
func ReleaseTaobaoFuwuPurchaseOrderConfirmAPIRequest(v *TaobaoFuwuPurchaseOrderConfirmAPIRequest) {
	v.Reset()
	poolTaobaoFuwuPurchaseOrderConfirmAPIRequest.Put(v)
}
