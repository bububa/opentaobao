package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单条订单跟踪详情 APIRequest
taobao.jds.trade.traces.get

获取聚石塔数据共享的交易全链路信息
*/
type TaobaoJdsTradeTracesGetRequest struct {
    model.Params

    // 淘宝的订单tid
    tid   int64 

}

func NewTaobaoJdsTradeTracesGetRequest() *TaobaoJdsTradeTracesGetRequest{
    return &TaobaoJdsTradeTracesGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJdsTradeTracesGetRequest) GetApiMethodName() string {
    return "taobao.jds.trade.traces.get"
}

func (r TaobaoJdsTradeTracesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJdsTradeTracesGetRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoJdsTradeTracesGetRequest) GetTid() int64 {
    return r.tid
}

