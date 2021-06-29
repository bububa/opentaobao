package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分单结果同步给星盘 API请求
taobao.omniorder.allocatedinfo.sync

ISV分单完成，将分单结果同步给星盘
*/
type TaobaoOmniorderAllocatedinfoSyncRequest struct {
    model.Params
    // 淘宝交易主订单ID
    _tid   int64
    // 分单状态，如： 等待中(Waiting)，已分单(Allocated)，分单失败(AllocateFail)
    _status   string
    // 分单结果消息, 如果status为AllocateFail, 则表示失败的理由.
    _message   string
    // 1231243213213
    _reportTimestamp   int64
    // 门店的分单列表
    _subOrderList   []StoreAllocatedResult
    // 跟踪Id
    _traceId   string
}

// 初始化TaobaoOmniorderAllocatedinfoSyncRequest对象
func NewTaobaoOmniorderAllocatedinfoSyncRequest() *TaobaoOmniorderAllocatedinfoSyncRequest{
    return &TaobaoOmniorderAllocatedinfoSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderAllocatedinfoSyncRequest) GetApiMethodName() string {
    return "taobao.omniorder.allocatedinfo.sync"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderAllocatedinfoSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 淘宝交易主订单ID
func (r *TaobaoOmniorderAllocatedinfoSyncRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoOmniorderAllocatedinfoSyncRequest) GetTid() int64 {
    return r._tid
}
// Status Setter
// 分单状态，如： 等待中(Waiting)，已分单(Allocated)，分单失败(AllocateFail)
func (r *TaobaoOmniorderAllocatedinfoSyncRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoOmniorderAllocatedinfoSyncRequest) GetStatus() string {
    return r._status
}
// Message Setter
// 分单结果消息, 如果status为AllocateFail, 则表示失败的理由.
func (r *TaobaoOmniorderAllocatedinfoSyncRequest) SetMessage(_message string) error {
    r._message = _message
    r.Set("message", _message)
    return nil
}

// Message Getter
func (r TaobaoOmniorderAllocatedinfoSyncRequest) GetMessage() string {
    return r._message
}
// ReportTimestamp Setter
// 1231243213213
func (r *TaobaoOmniorderAllocatedinfoSyncRequest) SetReportTimestamp(_reportTimestamp int64) error {
    r._reportTimestamp = _reportTimestamp
    r.Set("report_timestamp", _reportTimestamp)
    return nil
}

// ReportTimestamp Getter
func (r TaobaoOmniorderAllocatedinfoSyncRequest) GetReportTimestamp() int64 {
    return r._reportTimestamp
}
// SubOrderList Setter
// 门店的分单列表
func (r *TaobaoOmniorderAllocatedinfoSyncRequest) SetSubOrderList(_subOrderList []StoreAllocatedResult) error {
    r._subOrderList = _subOrderList
    r.Set("sub_order_list", _subOrderList)
    return nil
}

// SubOrderList Getter
func (r TaobaoOmniorderAllocatedinfoSyncRequest) GetSubOrderList() []StoreAllocatedResult {
    return r._subOrderList
}
// TraceId Setter
// 跟踪Id
func (r *TaobaoOmniorderAllocatedinfoSyncRequest) SetTraceId(_traceId string) error {
    r._traceId = _traceId
    r.Set("trace_id", _traceId)
    return nil
}

// TraceId Getter
func (r TaobaoOmniorderAllocatedinfoSyncRequest) GetTraceId() string {
    return r._traceId
}
