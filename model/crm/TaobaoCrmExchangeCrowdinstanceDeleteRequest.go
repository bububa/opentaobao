package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除人群实例中的指定买家 API请求
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

// 初始化TaobaoCrmExchangeCrowdinstanceDeleteRequest对象
func NewTaobaoCrmExchangeCrowdinstanceDeleteRequest() *TaobaoCrmExchangeCrowdinstanceDeleteRequest{
    return &TaobaoCrmExchangeCrowdinstanceDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCrmExchangeCrowdinstanceDeleteRequest) GetApiMethodName() string {
    return "taobao.crm.exchange.crowdinstance.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCrmExchangeCrowdinstanceDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Reason Setter
// 操作原因
func (r *TaobaoCrmExchangeCrowdinstanceDeleteRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

// Reason Getter
func (r TaobaoCrmExchangeCrowdinstanceDeleteRequest) GetReason() string {
    return r.reason
}
// CrowdInstanceId Setter
// 人群实例ID
func (r *TaobaoCrmExchangeCrowdinstanceDeleteRequest) SetCrowdInstanceId(crowdInstanceId int64) error {
    r.crowdInstanceId = crowdInstanceId
    r.Set("crowd_instance_id", crowdInstanceId)
    return nil
}

// CrowdInstanceId Getter
func (r TaobaoCrmExchangeCrowdinstanceDeleteRequest) GetCrowdInstanceId() int64 {
    return r.crowdInstanceId
}
// BuyerNick Setter
// 买家昵称
func (r *TaobaoCrmExchangeCrowdinstanceDeleteRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoCrmExchangeCrowdinstanceDeleteRequest) GetBuyerNick() string {
    return r.buyerNick
}
