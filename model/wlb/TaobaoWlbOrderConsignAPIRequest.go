package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbOrderConsignAPIRequest
物流宝订单已发货通知接口 API请求
taobao.wlb.order.consign

如果erp导入淘宝交易订单到物流宝，当物流宝订单已发货的时候，erp需要调用该接口来通知物流订单和淘宝交易订单已发货 */
type TaobaoWlbOrderConsignAPIRequest struct {
	model.Params
	// 物流宝订单编号
	_wlbOrderCode string
}

// NewTaobaoWlbOrderConsignRequest 初始化TaobaoWlbOrderConsignAPIRequest对象
func NewTaobaoWlbOrderConsignRequest() *TaobaoWlbOrderConsignAPIRequest {
	return &TaobaoWlbOrderConsignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbOrderConsignAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.order.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbOrderConsignAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is WlbOrderCode Setter
// 物流宝订单编号
func (r *TaobaoWlbOrderConsignAPIRequest) SetWlbOrderCode(_wlbOrderCode string) error {
	r._wlbOrderCode = _wlbOrderCode
	r.Set("wlb_order_code", _wlbOrderCode)
	return nil
}

// Get WlbOrderCode Getter
func (r TaobaoWlbOrderConsignAPIRequest) GetWlbOrderCode() string {
	return r._wlbOrderCode
}
