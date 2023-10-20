package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniorderallocatedinfosyncAPIRequest 分单结果同步给星盘 API请求
// taobao.omniorder.allocatedinfo.sync
//
// ISV分单完成，将分单结果同步给星盘
type TaobaoomniorderallocatedinfosyncAPIRequest struct {
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

// NewTaobaoomniorderallocatedinfosyncRequest 初始化TaobaoomniorderallocatedinfosyncAPIRequest对象
func NewTaobaoomniorderallocatedinfosyncRequest() *TaobaoomniorderallocatedinfosyncAPIRequest {
	return &TaobaoomniorderallocatedinfosyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomniorderallocatedinfosyncAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.allocatedinfo.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomniorderallocatedinfosyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomniorderallocatedinfosyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubOrderList is SubOrderList Setter
// 门店的分单列表
func (r *TaobaoomniorderallocatedinfosyncAPIRequest) SetSubOrderList(_subOrderList []StoreAllocatedResult) error {
	r._subOrderList = _subOrderList
	r.Set("sub_order_list", _subOrderList)
	return nil
}

// GetSubOrderList SubOrderList Getter
func (r TaobaoomniorderallocatedinfosyncAPIRequest) GetSubOrderList() []StoreAllocatedResult {
	return r._subOrderList
}

// SetTraceId is TraceId Setter
// 跟踪Id
func (r *TaobaoomniorderallocatedinfosyncAPIRequest) SetTraceId(_traceId string) error {
	r._traceId = _traceId
	r.Set("trace_id", _traceId)
	return nil
}

// GetTraceId TraceId Getter
func (r TaobaoomniorderallocatedinfosyncAPIRequest) GetTraceId() string {
	return r._traceId
}

// SetStatus is Status Setter
// 分单状态，如： 等待中(Waiting)，已分单(Allocated)，分单失败(AllocateFail)
func (r *TaobaoomniorderallocatedinfosyncAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoomniorderallocatedinfosyncAPIRequest) GetStatus() string {
	return r._status
}

// SetMessage is Message Setter
// 分单结果消息, 如果status为AllocateFail, 则表示失败的理由.
func (r *TaobaoomniorderallocatedinfosyncAPIRequest) SetMessage(_message string) error {
	r._message = _message
	r.Set("message", _message)
	return nil
}

// GetMessage Message Getter
func (r TaobaoomniorderallocatedinfosyncAPIRequest) GetMessage() string {
	return r._message
}

// SetTid is Tid Setter
// 淘宝交易主订单ID
func (r *TaobaoomniorderallocatedinfosyncAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoomniorderallocatedinfosyncAPIRequest) GetTid() int64 {
	return r._tid
}

// SetReportTimestamp is ReportTimestamp Setter
// 1231243213213
func (r *TaobaoomniorderallocatedinfosyncAPIRequest) SetReportTimestamp(_reportTimestamp int64) error {
	r._reportTimestamp = _reportTimestamp
	r.Set("report_timestamp", _reportTimestamp)
	return nil
}

// GetReportTimestamp ReportTimestamp Getter
func (r TaobaoomniorderallocatedinfosyncAPIRequest) GetReportTimestamp() int64 {
	return r._reportTimestamp
}
