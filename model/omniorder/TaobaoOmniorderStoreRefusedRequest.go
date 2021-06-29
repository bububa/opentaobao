package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Pos端门店拒单 API请求
taobao.omniorder.store.refused

ISV Pos端门店拒单，通知星盘
*/
type TaobaoOmniorderStoreRefusedRequest struct {
    model.Params
    // 淘宝交易主订单ID
    _tid   int64
    // 子订单列表
    _subOrderList   []SubOrder
    // ISV的系统时间
    _reportTimestamp   int64
    // 跟踪Id
    _traceId   string
}

// 初始化TaobaoOmniorderStoreRefusedRequest对象
func NewTaobaoOmniorderStoreRefusedRequest() *TaobaoOmniorderStoreRefusedRequest{
    return &TaobaoOmniorderStoreRefusedRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreRefusedRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.refused"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreRefusedRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 淘宝交易主订单ID
func (r *TaobaoOmniorderStoreRefusedRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoOmniorderStoreRefusedRequest) GetTid() int64 {
    return r._tid
}
// SubOrderList Setter
// 子订单列表
func (r *TaobaoOmniorderStoreRefusedRequest) SetSubOrderList(_subOrderList []SubOrder) error {
    r._subOrderList = _subOrderList
    r.Set("sub_order_list", _subOrderList)
    return nil
}

// SubOrderList Getter
func (r TaobaoOmniorderStoreRefusedRequest) GetSubOrderList() []SubOrder {
    return r._subOrderList
}
// ReportTimestamp Setter
// ISV的系统时间
func (r *TaobaoOmniorderStoreRefusedRequest) SetReportTimestamp(_reportTimestamp int64) error {
    r._reportTimestamp = _reportTimestamp
    r.Set("report_timestamp", _reportTimestamp)
    return nil
}

// ReportTimestamp Getter
func (r TaobaoOmniorderStoreRefusedRequest) GetReportTimestamp() int64 {
    return r._reportTimestamp
}
// TraceId Setter
// 跟踪Id
func (r *TaobaoOmniorderStoreRefusedRequest) SetTraceId(_traceId string) error {
    r._traceId = _traceId
    r.Set("trace_id", _traceId)
    return nil
}

// TraceId Getter
func (r TaobaoOmniorderStoreRefusedRequest) GetTraceId() string {
    return r._traceId
}
