package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除人群实例中的指定买家 APIRequest
taobao.crm.exchange.crowdinstance.delete

删除人群实例中的指定买家
*/
type TaobaoCrmExchangeCrowdinstanceDeleteRequest struct {
    model.Params

    // 操作原因
    reason   string 

    // 人群实例ID
    crowdInstanceId   int64 

    // 买家昵称
    buyerNick   string 

}

func NewTaobaoCrmExchangeCrowdinstanceDeleteRequest() *TaobaoCrmExchangeCrowdinstanceDeleteRequest{
    return &TaobaoCrmExchangeCrowdinstanceDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCrmExchangeCrowdinstanceDeleteRequest) GetApiMethodName() string {
    return "taobao.crm.exchange.crowdinstance.delete"
}

func (r TaobaoCrmExchangeCrowdinstanceDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCrmExchangeCrowdinstanceDeleteRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

func (r TaobaoCrmExchangeCrowdinstanceDeleteRequest) GetReason() string {
    return r.reason
}

func (r *TaobaoCrmExchangeCrowdinstanceDeleteRequest) SetCrowdInstanceId(crowdInstanceId int64) error {
    r.crowdInstanceId = crowdInstanceId
    r.Set("crowd_instance_id", crowdInstanceId)
    return nil
}

func (r TaobaoCrmExchangeCrowdinstanceDeleteRequest) GetCrowdInstanceId() int64 {
    return r.crowdInstanceId
}

func (r *TaobaoCrmExchangeCrowdinstanceDeleteRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

func (r TaobaoCrmExchangeCrowdinstanceDeleteRequest) GetBuyerNick() string {
    return r.buyerNick
}

