package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreAccpetedAPIRequest Pos端门店接单接口 API请求
// taobao.omniorder.store.accpeted
//
// ISV Pos端门店接单，通知星盘
type TaobaoOmniorderStoreAccpetedAPIRequest struct {
	model.Params
	// 子订单列表
	_subOrderList []StoreAcceptedResult
	// 跟踪Id
	_traceId string
	// 淘宝交易主订单ID
	_tid int64
	// ISV系统上报时间
	_reportTimestamp int64
}

// NewTaobaoOmniorderStoreAccpetedRequest 初始化TaobaoOmniorderStoreAccpetedAPIRequest对象
func NewTaobaoOmniorderStoreAccpetedRequest() *TaobaoOmniorderStoreAccpetedAPIRequest {
	return &TaobaoOmniorderStoreAccpetedAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniorderStoreAccpetedAPIRequest) Reset() {
	r._subOrderList = r._subOrderList[:0]
	r._traceId = ""
	r._tid = 0
	r._reportTimestamp = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreAccpetedAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.accpeted"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreAccpetedAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniorderStoreAccpetedAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubOrderList is SubOrderList Setter
// 子订单列表
func (r *TaobaoOmniorderStoreAccpetedAPIRequest) SetSubOrderList(_subOrderList []StoreAcceptedResult) error {
	r._subOrderList = _subOrderList
	r.Set("sub_order_list", _subOrderList)
	return nil
}

// GetSubOrderList SubOrderList Getter
func (r TaobaoOmniorderStoreAccpetedAPIRequest) GetSubOrderList() []StoreAcceptedResult {
	return r._subOrderList
}

// SetTraceId is TraceId Setter
// 跟踪Id
func (r *TaobaoOmniorderStoreAccpetedAPIRequest) SetTraceId(_traceId string) error {
	r._traceId = _traceId
	r.Set("trace_id", _traceId)
	return nil
}

// GetTraceId TraceId Getter
func (r TaobaoOmniorderStoreAccpetedAPIRequest) GetTraceId() string {
	return r._traceId
}

// SetTid is Tid Setter
// 淘宝交易主订单ID
func (r *TaobaoOmniorderStoreAccpetedAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoOmniorderStoreAccpetedAPIRequest) GetTid() int64 {
	return r._tid
}

// SetReportTimestamp is ReportTimestamp Setter
// ISV系统上报时间
func (r *TaobaoOmniorderStoreAccpetedAPIRequest) SetReportTimestamp(_reportTimestamp int64) error {
	r._reportTimestamp = _reportTimestamp
	r.Set("report_timestamp", _reportTimestamp)
	return nil
}

// GetReportTimestamp ReportTimestamp Getter
func (r TaobaoOmniorderStoreAccpetedAPIRequest) GetReportTimestamp() int64 {
	return r._reportTimestamp
}

var poolTaobaoOmniorderStoreAccpetedAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniorderStoreAccpetedRequest()
	},
}

// GetTaobaoOmniorderStoreAccpetedRequest 从 sync.Pool 获取 TaobaoOmniorderStoreAccpetedAPIRequest
func GetTaobaoOmniorderStoreAccpetedAPIRequest() *TaobaoOmniorderStoreAccpetedAPIRequest {
	return poolTaobaoOmniorderStoreAccpetedAPIRequest.Get().(*TaobaoOmniorderStoreAccpetedAPIRequest)
}

// ReleaseTaobaoOmniorderStoreAccpetedAPIRequest 将 TaobaoOmniorderStoreAccpetedAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniorderStoreAccpetedAPIRequest(v *TaobaoOmniorderStoreAccpetedAPIRequest) {
	v.Reset()
	poolTaobaoOmniorderStoreAccpetedAPIRequest.Put(v)
}
