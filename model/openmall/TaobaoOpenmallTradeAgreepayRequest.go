package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
openmall订单支付 APIRequest
taobao.openmall.trade.agreepay

openmall订单支付
*/
type TaobaoOpenmallTradeAgreepayRequest struct {
    model.Params

    // 媒体渠道，代表分销者的身份，签约支付宝代扣的渠道商淘宝账号nick
    distributor   string 

    // 淘宝交易单号
    tid   int64 

}

func NewTaobaoOpenmallTradeAgreepayRequest() *TaobaoOpenmallTradeAgreepayRequest{
    return &TaobaoOpenmallTradeAgreepayRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenmallTradeAgreepayRequest) GetApiMethodName() string {
    return "taobao.openmall.trade.agreepay"
}

func (r TaobaoOpenmallTradeAgreepayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenmallTradeAgreepayRequest) SetDistributor(distributor string) error {
    r.distributor = distributor
    r.Set("distributor", distributor)
    return nil
}

func (r TaobaoOpenmallTradeAgreepayRequest) GetDistributor() string {
    return r.distributor
}

func (r *TaobaoOpenmallTradeAgreepayRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoOpenmallTradeAgreepayRequest) GetTid() int64 {
    return r.tid
}

