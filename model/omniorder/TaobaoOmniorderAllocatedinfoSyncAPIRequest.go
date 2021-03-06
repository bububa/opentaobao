package omniorder

import (
	"net/url"

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
	// 分单状态，如： 等待中(Waiting)，已分单(Allocated)，分单失败(AllocateFail)
	_status string
	// 分单结果消息, 如果status为AllocateFail, 则表示失败的理由.
	_message string
	// 跟踪Id
	_traceId string
	// 淘宝交易主订单ID
	_tid int64
	// 1231243213213
	_reportTimestamp int64
}

// NewTaobaoOmniorderAllocatedinfoSyncRequest 初始化TaobaoOmniorderAllocatedinfoSyncAPIRequest对象
func NewTaobaoOmniorderAllocatedinfoSyncRequest() *TaobaoOmniorderAllocatedinfoSyncAPIRequest {
	return &TaobaoOmniorderAllocatedinfoSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderAllocatedinfoSyncAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.allocatedinfo.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderAllocatedinfoSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
