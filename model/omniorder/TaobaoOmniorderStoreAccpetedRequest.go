package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Pos端门店接单接口 API请求
taobao.omniorder.store.accpeted

ISV Pos端门店接单，通知星盘
*/
type TaobaoOmniorderStoreAccpetedRequest struct {
    model.Params
    // 淘宝交易主订单ID
    tid   int64
    // 子订单列表
    subOrderList   []StoreAcceptedResult
    // ISV系统上报时间
    reportTimestamp   int64
    // 跟踪Id
    traceId   string
}

// 初始化TaobaoOmniorderStoreAccpetedRequest对象
func NewTaobaoOmniorderStoreAccpetedRequest() *TaobaoOmniorderStoreAccpetedRequest{
    return &TaobaoOmniorderStoreAccpetedRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreAccpetedRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.accpeted"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreAccpetedRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 淘宝交易主订单ID
func (r *TaobaoOmniorderStoreAccpetedRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoOmniorderStoreAccpetedRequest) GetTid() int64 {
    return r.tid
}
// SubOrderList Setter
// 子订单列表
func (r *TaobaoOmniorderStoreAccpetedRequest) SetSubOrderList(subOrderList []StoreAcceptedResult) error {
    r.subOrderList = subOrderList
    r.Set("sub_order_list", subOrderList)
    return nil
}

// SubOrderList Getter
func (r TaobaoOmniorderStoreAccpetedRequest) GetSubOrderList() []StoreAcceptedResult {
    return r.subOrderList
}
// ReportTimestamp Setter
// ISV系统上报时间
func (r *TaobaoOmniorderStoreAccpetedRequest) SetReportTimestamp(reportTimestamp int64) error {
    r.reportTimestamp = reportTimestamp
    r.Set("report_timestamp", reportTimestamp)
    return nil
}

// ReportTimestamp Getter
func (r TaobaoOmniorderStoreAccpetedRequest) GetReportTimestamp() int64 {
    return r.reportTimestamp
}
// TraceId Setter
// 跟踪Id
func (r *TaobaoOmniorderStoreAccpetedRequest) SetTraceId(traceId string) error {
    r.traceId = traceId
    r.Set("trace_id", traceId)
    return nil
}

// TraceId Getter
func (r TaobaoOmniorderStoreAccpetedRequest) GetTraceId() string {
    return r.traceId
}
