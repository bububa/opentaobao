package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认收货 APIRequest
taobao.openmall.trade.confirm

确认订单收货
*/
type TaobaoOpenmallTradeConfirmRequest struct {
    model.Params

    // 分销者信息
    distributor   string 

    // 淘宝订单号
    tid   int64 

}

func NewTaobaoOpenmallTradeConfirmRequest() *TaobaoOpenmallTradeConfirmRequest{
    return &TaobaoOpenmallTradeConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenmallTradeConfirmRequest) GetApiMethodName() string {
    return "taobao.openmall.trade.confirm"
}

func (r TaobaoOpenmallTradeConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenmallTradeConfirmRequest) SetDistributor(distributor string) error {
    r.distributor = distributor
    r.Set("distributor", distributor)
    return nil
}

func (r TaobaoOpenmallTradeConfirmRequest) GetDistributor() string {
    return r.distributor
}

func (r *TaobaoOpenmallTradeConfirmRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoOpenmallTradeConfirmRequest) GetTid() int64 {
    return r.tid
}

