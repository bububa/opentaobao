package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流流转信息查询 APIRequest
taobao.logistics.trace.search

用户根据淘宝交易号查询物流流转信息，如2010-8-10 15：23：00到达杭州集散地。<br/>此接口的返回信息都由物流公司提供。（备注：使用线下发货（offline.send）的运单，不支持运单状态的实时跟踪，只要一发货，状态就会变为<status>对方已签收</status>，该字段仅对线上发货（online.send）的运单有效。）
*/
type TaobaoLogisticsTraceSearchRequest struct {
    model.Params

    // 淘宝交易号，请勿传非淘宝交易号
    tid   int64 

    // 表明是否是拆单，默认值0，1表示拆单
    isSplit   int64 

    // 拆单子订单列表，当is_split=1时，需要传人；对应的数据是：子订单号的列表。
    subTid   []Number 

}

func NewTaobaoLogisticsTraceSearchRequest() *TaobaoLogisticsTraceSearchRequest{
    return &TaobaoLogisticsTraceSearchRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoLogisticsTraceSearchRequest) GetApiMethodName() string {
    return "taobao.logistics.trace.search"
}

func (r TaobaoLogisticsTraceSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoLogisticsTraceSearchRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoLogisticsTraceSearchRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoLogisticsTraceSearchRequest) SetIsSplit(isSplit int64) error {
    r.isSplit = isSplit
    r.Set("is_split", isSplit)
    return nil
}

func (r TaobaoLogisticsTraceSearchRequest) GetIsSplit() int64 {
    return r.isSplit
}

func (r *TaobaoLogisticsTraceSearchRequest) SetSubTid(subTid []Number) error {
    r.subTid = subTid
    r.Set("sub_tid", subTid)
    return nil
}

func (r TaobaoLogisticsTraceSearchRequest) GetSubTid() []Number {
    return r.subTid
}

