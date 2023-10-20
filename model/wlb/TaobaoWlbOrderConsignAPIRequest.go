package wlb

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbOrderConsignAPIRequest 物流宝订单已发货通知接口 API请求
// taobao.wlb.order.consign
//
// 如果erp导入淘宝交易订单到物流宝，当物流宝订单已发货的时候，erp需要调用该接口来通知物流订单和淘宝交易订单已发货
type TaobaoWlbOrderConsignAPIRequest struct {
	model.Params
	// 物流宝订单编号
	_wlbOrderCode string
}

// NewTaobaoWlbOrderConsignRequest 初始化TaobaoWlbOrderConsignAPIRequest对象
func NewTaobaoWlbOrderConsignRequest() *TaobaoWlbOrderConsignAPIRequest {
	return &TaobaoWlbOrderConsignAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbOrderConsignAPIRequest) Reset() {
	r._wlbOrderCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbOrderConsignAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.order.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbOrderConsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbOrderConsignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWlbOrderCode is WlbOrderCode Setter
// 物流宝订单编号
func (r *TaobaoWlbOrderConsignAPIRequest) SetWlbOrderCode(_wlbOrderCode string) error {
	r._wlbOrderCode = _wlbOrderCode
	r.Set("wlb_order_code", _wlbOrderCode)
	return nil
}

// GetWlbOrderCode WlbOrderCode Getter
func (r TaobaoWlbOrderConsignAPIRequest) GetWlbOrderCode() string {
	return r._wlbOrderCode
}

var poolTaobaoWlbOrderConsignAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbOrderConsignRequest()
	},
}

// GetTaobaoWlbOrderConsignRequest 从 sync.Pool 获取 TaobaoWlbOrderConsignAPIRequest
func GetTaobaoWlbOrderConsignAPIRequest() *TaobaoWlbOrderConsignAPIRequest {
	return poolTaobaoWlbOrderConsignAPIRequest.Get().(*TaobaoWlbOrderConsignAPIRequest)
}

// ReleaseTaobaoWlbOrderConsignAPIRequest 将 TaobaoWlbOrderConsignAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbOrderConsignAPIRequest(v *TaobaoWlbOrderConsignAPIRequest) {
	v.Reset()
	poolTaobaoWlbOrderConsignAPIRequest.Put(v)
}
