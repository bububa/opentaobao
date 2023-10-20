package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderAllocatedinfoSyncAPIRequest 分单结果同步给星盘 API请求
// taobao.omniorder.allocatedinfo.sync
//
// ISV分单完成，将分单结果同步给星盘
type TaobaoOmniorderAllocatedinfoSyncAPIRequest struct {
	model.Params
	// 门店的分单列表
	_subOrderList []StoreAllocatedResult
	// 跟踪Id
	_traceId string
	// 分单状态，如： 等待中(Waiting)，已分单(Allocated)，分单失败(AllocateFail)
	_status string
	// 分单结果消息, 如果status为AllocateFail, 则表示失败的理由.
	_message string
	// 淘宝交易主订单ID
	_tid int64
	// 1231243213213
	_reportTimestamp int64
}

// NewTaobaoOmniorderAllocatedinfoSyncRequest 初始化TaobaoOmniorderAllocatedinfoSyncAPIRequest对象
func NewTaobaoOmniorderAllocatedinfoSyncRequest() *TaobaoOmniorderAllocatedinfoSyncAPIRequest {
	return &TaobaoOmniorderAllocatedinfoSyncAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniorderAllocatedinfoSyncAPIRequest) Reset() {
	r._subOrderList = r._subOrderList[:0]
	r._traceId = ""
	r._status = ""
	r._message = ""
	r._tid = 0
	r._reportTimestamp = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderAllocatedinfoSyncAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.allocatedinfo.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderAllocatedinfoSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniorderAllocatedinfoSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubOrderList is SubOrderList Setter
// 门店的分单列表
func (r *TaobaoOmniorderAllocatedinfoSyncAPIRequest) SetSubOrderList(_subOrderList []StoreAllocatedResult) error {
	r._subOrderList = _subOrderList
	r.Set("sub_order_list", _subOrderList)
	return nil
}

// GetSubOrderList SubOrderList Getter
func (r TaobaoOmniorderAllocatedinfoSyncAPIRequest) GetSubOrderList() []StoreAllocatedResult {
	return r._subOrderList
}

// SetTraceId is TraceId Setter
// 跟踪Id
func (r *TaobaoOmniorderAllocatedinfoSyncAPIRequest) SetTraceId(_traceId string) error {
	r._traceId = _traceId
	r.Set("trace_id", _traceId)
	return nil
}

// GetTraceId TraceId Getter
func (r TaobaoOmniorderAllocatedinfoSyncAPIRequest) GetTraceId() string {
	return r._traceId
}

// SetStatus is Status Setter
// 分单状态，如： 等待中(Waiting)，已分单(Allocated)，分单失败(AllocateFail)
func (r *TaobaoOmniorderAllocatedinfoSyncAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoOmniorderAllocatedinfoSyncAPIRequest) GetStatus() string {
	return r._status
}

// SetMessage is Message Setter
// 分单结果消息, 如果status为AllocateFail, 则表示失败的理由.
func (r *TaobaoOmniorderAllocatedinfoSyncAPIRequest) SetMessage(_message string) error {
	r._message = _message
	r.Set("message", _message)
	return nil
}

// GetMessage Message Getter
func (r TaobaoOmniorderAllocatedinfoSyncAPIRequest) GetMessage() string {
	return r._message
}

// SetTid is Tid Setter
// 淘宝交易主订单ID
func (r *TaobaoOmniorderAllocatedinfoSyncAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoOmniorderAllocatedinfoSyncAPIRequest) GetTid() int64 {
	return r._tid
}

// SetReportTimestamp is ReportTimestamp Setter
// 1231243213213
func (r *TaobaoOmniorderAllocatedinfoSyncAPIRequest) SetReportTimestamp(_reportTimestamp int64) error {
	r._reportTimestamp = _reportTimestamp
	r.Set("report_timestamp", _reportTimestamp)
	return nil
}

// GetReportTimestamp ReportTimestamp Getter
func (r TaobaoOmniorderAllocatedinfoSyncAPIRequest) GetReportTimestamp() int64 {
	return r._reportTimestamp
}

var poolTaobaoOmniorderAllocatedinfoSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniorderAllocatedinfoSyncRequest()
	},
}

// GetTaobaoOmniorderAllocatedinfoSyncRequest 从 sync.Pool 获取 TaobaoOmniorderAllocatedinfoSyncAPIRequest
func GetTaobaoOmniorderAllocatedinfoSyncAPIRequest() *TaobaoOmniorderAllocatedinfoSyncAPIRequest {
	return poolTaobaoOmniorderAllocatedinfoSyncAPIRequest.Get().(*TaobaoOmniorderAllocatedinfoSyncAPIRequest)
}

// ReleaseTaobaoOmniorderAllocatedinfoSyncAPIRequest 将 TaobaoOmniorderAllocatedinfoSyncAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniorderAllocatedinfoSyncAPIRequest(v *TaobaoOmniorderAllocatedinfoSyncAPIRequest) {
	v.Reset()
	poolTaobaoOmniorderAllocatedinfoSyncAPIRequest.Put(v)
}
