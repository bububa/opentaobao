package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlborderconsignAPIRequest 物流宝订单已发货通知接口 API请求
// taobao.wlb.order.consign
//
// 如果erp导入淘宝交易订单到物流宝，当物流宝订单已发货的时候，erp需要调用该接口来通知物流订单和淘宝交易订单已发货
type TaobaowlborderconsignAPIRequest struct {
	model.Params
	// 物流宝订单编号
	_wlbOrderCode string
}

// NewTaobaowlborderconsignRequest 初始化TaobaowlborderconsignAPIRequest对象
func NewTaobaowlborderconsignRequest() *TaobaowlborderconsignAPIRequest {
	return &TaobaowlborderconsignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlborderconsignAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.order.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlborderconsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlborderconsignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWlbOrderCode is WlbOrderCode Setter
// 物流宝订单编号
func (r *TaobaowlborderconsignAPIRequest) SetWlbOrderCode(_wlbOrderCode string) error {
	r._wlbOrderCode = _wlbOrderCode
	r.Set("wlb_order_code", _wlbOrderCode)
	return nil
}

// GetWlbOrderCode WlbOrderCode Getter
func (r TaobaowlborderconsignAPIRequest) GetWlbOrderCode() string {
	return r._wlbOrderCode
}
