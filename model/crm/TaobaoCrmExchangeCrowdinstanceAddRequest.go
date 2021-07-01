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
type TaobaoCrmExchangeCrowdinstanceAddAPIRequest struct {
    model.Params
    // 操作原因
    _reason   string
    // 人群实例ID
    _crowdInstanceId   int64
    // 买家昵称
    _buyerNick   string
}

// 初始化TaobaoCrmExchangeCrowdinstanceAddAPIRequest对象
func NewTaobaoCrmExchangeCrowdinstanceAddRequest() *TaobaoCrmExchangeCrowdinstanceAddAPIRequest{
    return &TaobaoCrmExchangeCrowdinstanceAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCrmExchangeCrowdinstanceAddAPIRequest) GetApiMethodName() string {
    return "taobao.crm.exchange.crowdinstance.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCrmExchangeCrowdinstanceAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Reason Setter
// 操作原因
func (r *TaobaoCrmExchangeCrowdinstanceAddAPIRequest) SetReason(_reason string) error {
    r._reason = _reason
    r.Set("reason", _reason)
    return nil
}

// Reason Getter
func (r TaobaoCrmExchangeCrowdinstanceAddAPIRequest) GetReason() string {
    return r._reason
}
// CrowdInstanceId Setter
// 人群实例ID
func (r *TaobaoCrmExchangeCrowdinstanceAddAPIRequest) SetCrowdInstanceId(_crowdInstanceId int64) error {
    r._crowdInstanceId = _crowdInstanceId
    r.Set("crowd_instance_id", _crowdInstanceId)
    return nil
}

// CrowdInstanceId Getter
func (r TaobaoCrmExchangeCrowdinstanceAddAPIRequest) GetCrowdInstanceId() int64 {
    return r._crowdInstanceId
}
// BuyerNick Setter
// 买家昵称
func (r *TaobaoCrmExchangeCrowdinstanceAddAPIRequest) SetBuyerNick(_buyerNick string) error {
    r._buyerNick = _buyerNick
    r.Set("buyer_nick", _buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoCrmExchangeCrowdinstanceAddAPIRequest) GetBuyerNick() string {
    return r._buyerNick
}
