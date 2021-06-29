package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
向活动人群实例中增加买家 API请求
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

// 初始化TaobaoCrmExchangeCrowdinstanceAddRequest对象
func NewTaobaoCrmExchangeCrowdinstanceAddRequest() *TaobaoCrmExchangeCrowdinstanceAddRequest{
    return &TaobaoCrmExchangeCrowdinstanceAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCrmExchangeCrowdinstanceAddRequest) GetApiMethodName() string {
    return "taobao.crm.exchange.crowdinstance.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCrmExchangeCrowdinstanceAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Reason Setter
// 操作原因
func (r *TaobaoCrmExchangeCrowdinstanceAddRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

// Reason Getter
func (r TaobaoCrmExchangeCrowdinstanceAddRequest) GetReason() string {
    return r.reason
}
// CrowdInstanceId Setter
// 人群实例ID
func (r *TaobaoCrmExchangeCrowdinstanceAddRequest) SetCrowdInstanceId(crowdInstanceId int64) error {
    r.crowdInstanceId = crowdInstanceId
    r.Set("crowd_instance_id", crowdInstanceId)
    return nil
}

// CrowdInstanceId Getter
func (r TaobaoCrmExchangeCrowdinstanceAddRequest) GetCrowdInstanceId() int64 {
    return r.crowdInstanceId
}
// BuyerNick Setter
// 买家昵称
func (r *TaobaoCrmExchangeCrowdinstanceAddRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoCrmExchangeCrowdinstanceAddRequest) GetBuyerNick() string {
    return r.buyerNick
}
