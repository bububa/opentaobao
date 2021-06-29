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
    _reason   string
    // 人群实例ID
    _crowdInstanceId   int64
    // 买家昵称
    _buyerNick   string
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
func (r *TaobaoCrmExchangeCrowdinstanceAddRequest) SetReason(_reason string) error {
    r._reason = _reason
    r.Set("reason", _reason)
    return nil
}

// Reason Getter
func (r TaobaoCrmExchangeCrowdinstanceAddRequest) GetReason() string {
    return r._reason
}
// CrowdInstanceId Setter
// 人群实例ID
func (r *TaobaoCrmExchangeCrowdinstanceAddRequest) SetCrowdInstanceId(_crowdInstanceId int64) error {
    r._crowdInstanceId = _crowdInstanceId
    r.Set("crowd_instance_id", _crowdInstanceId)
    return nil
}

// CrowdInstanceId Getter
func (r TaobaoCrmExchangeCrowdinstanceAddRequest) GetCrowdInstanceId() int64 {
    return r._crowdInstanceId
}
// BuyerNick Setter
// 买家昵称
func (r *TaobaoCrmExchangeCrowdinstanceAddRequest) SetBuyerNick(_buyerNick string) error {
    r._buyerNick = _buyerNick
    r.Set("buyer_nick", _buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoCrmExchangeCrowdinstanceAddRequest) GetBuyerNick() string {
    return r._buyerNick
}
