package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
向活动人群实例中增加买家 APIRequest
taobao.crm.exchange.crowdinstance.add

向活动人群实例中增加买家
*/
type TaobaoCrmExchangeCrowdinstanceAddRequest struct {
    model.Params

    // 操作原因
    reason   string 

    // 人群实例ID
    crowdInstanceId   int64 

    // 买家昵称
    buyerNick   string 

}

func NewTaobaoCrmExchangeCrowdinstanceAddRequest() *TaobaoCrmExchangeCrowdinstanceAddRequest{
    return &TaobaoCrmExchangeCrowdinstanceAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCrmExchangeCrowdinstanceAddRequest) GetApiMethodName() string {
    return "taobao.crm.exchange.crowdinstance.add"
}

func (r TaobaoCrmExchangeCrowdinstanceAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCrmExchangeCrowdinstanceAddRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

func (r TaobaoCrmExchangeCrowdinstanceAddRequest) GetReason() string {
    return r.reason
}

func (r *TaobaoCrmExchangeCrowdinstanceAddRequest) SetCrowdInstanceId(crowdInstanceId int64) error {
    r.crowdInstanceId = crowdInstanceId
    r.Set("crowd_instance_id", crowdInstanceId)
    return nil
}

func (r TaobaoCrmExchangeCrowdinstanceAddRequest) GetCrowdInstanceId() int64 {
    return r.crowdInstanceId
}

func (r *TaobaoCrmExchangeCrowdinstanceAddRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

func (r TaobaoCrmExchangeCrowdinstanceAddRequest) GetBuyerNick() string {
    return r.buyerNick
}

