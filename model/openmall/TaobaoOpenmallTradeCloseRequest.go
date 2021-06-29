package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关闭订单 APIRequest
taobao.openmall.trade.close

关闭订单
*/
type TaobaoOpenmallTradeCloseRequest struct {
    model.Params

    // 分销者信息
    distributor   string 

    // 关单原因
    reason   string 

    // 淘宝订单号
    tid   int64 

}

func NewTaobaoOpenmallTradeCloseRequest() *TaobaoOpenmallTradeCloseRequest{
    return &TaobaoOpenmallTradeCloseRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenmallTradeCloseRequest) GetApiMethodName() string {
    return "taobao.openmall.trade.close"
}

func (r TaobaoOpenmallTradeCloseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenmallTradeCloseRequest) SetDistributor(distributor string) error {
    r.distributor = distributor
    r.Set("distributor", distributor)
    return nil
}

func (r TaobaoOpenmallTradeCloseRequest) GetDistributor() string {
    return r.distributor
}

func (r *TaobaoOpenmallTradeCloseRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

func (r TaobaoOpenmallTradeCloseRequest) GetReason() string {
    return r.reason
}

func (r *TaobaoOpenmallTradeCloseRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoOpenmallTradeCloseRequest) GetTid() int64 {
    return r.tid
}

