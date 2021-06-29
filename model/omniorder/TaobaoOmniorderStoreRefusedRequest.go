package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Pos端门店拒单 APIRequest
taobao.omniorder.store.refused

ISV Pos端门店拒单，通知星盘
*/
type TaobaoOmniorderStoreRefusedRequest struct {
    model.Params

    // 淘宝交易主订单ID
    tid   int64 

    // 子订单列表
    subOrderList   []SubOrder 

    // ISV的系统时间
    reportTimestamp   int64 

    // 跟踪Id
    traceId   string 

}

func NewTaobaoOmniorderStoreRefusedRequest() *TaobaoOmniorderStoreRefusedRequest{
    return &TaobaoOmniorderStoreRefusedRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniorderStoreRefusedRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.refused"
}

func (r TaobaoOmniorderStoreRefusedRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniorderStoreRefusedRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoOmniorderStoreRefusedRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoOmniorderStoreRefusedRequest) SetSubOrderList(subOrderList []SubOrder) error {
    r.subOrderList = subOrderList
    r.Set("sub_order_list", subOrderList)
    return nil
}

func (r TaobaoOmniorderStoreRefusedRequest) GetSubOrderList() []SubOrder {
    return r.subOrderList
}

func (r *TaobaoOmniorderStoreRefusedRequest) SetReportTimestamp(reportTimestamp int64) error {
    r.reportTimestamp = reportTimestamp
    r.Set("report_timestamp", reportTimestamp)
    return nil
}

func (r TaobaoOmniorderStoreRefusedRequest) GetReportTimestamp() int64 {
    return r.reportTimestamp
}

func (r *TaobaoOmniorderStoreRefusedRequest) SetTraceId(traceId string) error {
    r.traceId = traceId
    r.Set("trace_id", traceId)
    return nil
}

func (r TaobaoOmniorderStoreRefusedRequest) GetTraceId() string {
    return r.traceId
}

