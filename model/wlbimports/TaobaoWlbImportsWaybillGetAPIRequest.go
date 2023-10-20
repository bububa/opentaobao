package wlbimports

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbImportsWaybillGetAPIRequest 获取运单信息【云打印】 API请求
// taobao.wlb.imports.waybill.get
//
// 一般进口商家，获取订单的电子面单链接地址
type TaobaoWlbImportsWaybillGetAPIRequest struct {
	model.Params
	// 物流订单号
	_orderCode string
}

// NewTaobaoWlbImportsWaybillGetRequest 初始化TaobaoWlbImportsWaybillGetAPIRequest对象
func NewTaobaoWlbImportsWaybillGetRequest() *TaobaoWlbImportsWaybillGetAPIRequest {
	return &TaobaoWlbImportsWaybillGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbImportsWaybillGetAPIRequest) Reset() {
	r._orderCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbImportsWaybillGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.imports.waybill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbImportsWaybillGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbImportsWaybillGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCode is OrderCode Setter
// 物流订单号
func (r *TaobaoWlbImportsWaybillGetAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r TaobaoWlbImportsWaybillGetAPIRequest) GetOrderCode() string {
	return r._orderCode
}

var poolTaobaoWlbImportsWaybillGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbImportsWaybillGetRequest()
	},
}

// GetTaobaoWlbImportsWaybillGetRequest 从 sync.Pool 获取 TaobaoWlbImportsWaybillGetAPIRequest
func GetTaobaoWlbImportsWaybillGetAPIRequest() *TaobaoWlbImportsWaybillGetAPIRequest {
	return poolTaobaoWlbImportsWaybillGetAPIRequest.Get().(*TaobaoWlbImportsWaybillGetAPIRequest)
}

// ReleaseTaobaoWlbImportsWaybillGetAPIRequest 将 TaobaoWlbImportsWaybillGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbImportsWaybillGetAPIRequest(v *TaobaoWlbImportsWaybillGetAPIRequest) {
	v.Reset()
	poolTaobaoWlbImportsWaybillGetAPIRequest.Put(v)
}
