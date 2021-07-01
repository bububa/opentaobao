package jstsecret

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取订单消费者的隐私号码 API请求
taobao.jst.secret.get

根据订单号获取消费者的隐私号
*/
type TaobaoJstSecretGetAPIRequest struct {
    model.Params
    // 订单号
    _tid   int64
    // 隐私号类型，1=手机号，2=分机号，默认为1；分机号使用时拨打手机号转分机号
    _type   int64
    // 隐私号码过期天数，默认30天，取值范围[1,30]
    _expireDays   int64
}

// 初始化TaobaoJstSecretGetAPIRequest对象
func NewTaobaoJstSecretGetRequest() *TaobaoJstSecretGetAPIRequest{
    return &TaobaoJstSecretGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstSecretGetAPIRequest) GetApiMethodName() string {
    return "taobao.jst.secret.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstSecretGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 订单号
func (r *TaobaoJstSecretGetAPIRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoJstSecretGetAPIRequest) GetTid() int64 {
    return r._tid
}
// Type Setter
// 隐私号类型，1=手机号，2=分机号，默认为1；分机号使用时拨打手机号转分机号
func (r *TaobaoJstSecretGetAPIRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoJstSecretGetAPIRequest) GetType() int64 {
    return r._type
}
// ExpireDays Setter
// 隐私号码过期天数，默认30天，取值范围[1,30]
func (r *TaobaoJstSecretGetAPIRequest) SetExpireDays(_expireDays int64) error {
    r._expireDays = _expireDays
    r.Set("expire_days", _expireDays)
    return nil
}

// ExpireDays Getter
func (r TaobaoJstSecretGetAPIRequest) GetExpireDays() int64 {
    return r._expireDays
}
