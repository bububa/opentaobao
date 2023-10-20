package waybill

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWaybillIQuerydetailAPIRequest 查面单号状态v1.0 API请求
// taobao.wlb.waybill.i.querydetail
//
// 查看面单号的当前状态，如签收、发货、失效等。
type TaobaoWlbWaybillIQuerydetailAPIRequest struct {
	model.Params
	// 面单查询请求
	_waybillDetailQueryRequest *WaybillDetailQueryRequest
}

// NewTaobaoWlbWaybillIQuerydetailRequest 初始化TaobaoWlbWaybillIQuerydetailAPIRequest对象
func NewTaobaoWlbWaybillIQuerydetailRequest() *TaobaoWlbWaybillIQuerydetailAPIRequest {
	return &TaobaoWlbWaybillIQuerydetailAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbWaybillIQuerydetailAPIRequest) Reset() {
	r._waybillDetailQueryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbWaybillIQuerydetailAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.i.querydetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbWaybillIQuerydetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbWaybillIQuerydetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWaybillDetailQueryRequest is WaybillDetailQueryRequest Setter
// 面单查询请求
func (r *TaobaoWlbWaybillIQuerydetailAPIRequest) SetWaybillDetailQueryRequest(_waybillDetailQueryRequest *WaybillDetailQueryRequest) error {
	r._waybillDetailQueryRequest = _waybillDetailQueryRequest
	r.Set("waybill_detail_query_request", _waybillDetailQueryRequest)
	return nil
}

// GetWaybillDetailQueryRequest WaybillDetailQueryRequest Getter
func (r TaobaoWlbWaybillIQuerydetailAPIRequest) GetWaybillDetailQueryRequest() *WaybillDetailQueryRequest {
	return r._waybillDetailQueryRequest
}

var poolTaobaoWlbWaybillIQuerydetailAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbWaybillIQuerydetailRequest()
	},
}

// GetTaobaoWlbWaybillIQuerydetailRequest 从 sync.Pool 获取 TaobaoWlbWaybillIQuerydetailAPIRequest
func GetTaobaoWlbWaybillIQuerydetailAPIRequest() *TaobaoWlbWaybillIQuerydetailAPIRequest {
	return poolTaobaoWlbWaybillIQuerydetailAPIRequest.Get().(*TaobaoWlbWaybillIQuerydetailAPIRequest)
}

// ReleaseTaobaoWlbWaybillIQuerydetailAPIRequest 将 TaobaoWlbWaybillIQuerydetailAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbWaybillIQuerydetailAPIRequest(v *TaobaoWlbWaybillIQuerydetailAPIRequest) {
	v.Reset()
	poolTaobaoWlbWaybillIQuerydetailAPIRequest.Put(v)
}
