package wlbimports

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbImportsOrderCancelAPIRequest 一般进口取消物流订单 API请求
// taobao.wlb.imports.order.cancel
//
// 商家在发货后，需要对订单进行取消，如果仓库已经收货则无法取消。
type TaobaoWlbImportsOrderCancelAPIRequest struct {
	model.Params
	// 物流订单编号
	_lgorderCode string
}

// NewTaobaoWlbImportsOrderCancelRequest 初始化TaobaoWlbImportsOrderCancelAPIRequest对象
func NewTaobaoWlbImportsOrderCancelRequest() *TaobaoWlbImportsOrderCancelAPIRequest {
	return &TaobaoWlbImportsOrderCancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbImportsOrderCancelAPIRequest) Reset() {
	r._lgorderCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbImportsOrderCancelAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.imports.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbImportsOrderCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbImportsOrderCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLgorderCode is LgorderCode Setter
// 物流订单编号
func (r *TaobaoWlbImportsOrderCancelAPIRequest) SetLgorderCode(_lgorderCode string) error {
	r._lgorderCode = _lgorderCode
	r.Set("lgorder_code", _lgorderCode)
	return nil
}

// GetLgorderCode LgorderCode Getter
func (r TaobaoWlbImportsOrderCancelAPIRequest) GetLgorderCode() string {
	return r._lgorderCode
}

var poolTaobaoWlbImportsOrderCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbImportsOrderCancelRequest()
	},
}

// GetTaobaoWlbImportsOrderCancelRequest 从 sync.Pool 获取 TaobaoWlbImportsOrderCancelAPIRequest
func GetTaobaoWlbImportsOrderCancelAPIRequest() *TaobaoWlbImportsOrderCancelAPIRequest {
	return poolTaobaoWlbImportsOrderCancelAPIRequest.Get().(*TaobaoWlbImportsOrderCancelAPIRequest)
}

// ReleaseTaobaoWlbImportsOrderCancelAPIRequest 将 TaobaoWlbImportsOrderCancelAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbImportsOrderCancelAPIRequest(v *TaobaoWlbImportsOrderCancelAPIRequest) {
	v.Reset()
	poolTaobaoWlbImportsOrderCancelAPIRequest.Put(v)
}
