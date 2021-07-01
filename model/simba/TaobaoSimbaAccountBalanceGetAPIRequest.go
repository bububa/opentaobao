package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取实时余额，”元”为单位 API请求
taobao.simba.account.balance.get

获取实时余额，”元”为单位
*/
type TaobaoSimbaAccountBalanceGetAPIRequest struct {
    model.Params
    // 主人昵称
    _nick   string
}

// 初始化TaobaoSimbaAccountBalanceGetAPIRequest对象
func NewTaobaoSimbaAccountBalanceGetRequest() *TaobaoSimbaAccountBalanceGetAPIRequest{
    return &TaobaoSimbaAccountBalanceGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAccountBalanceGetAPIRequest) GetApiMethodName() string {
    return "taobao.simba.account.balance.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAccountBalanceGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaAccountBalanceGetAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaAccountBalanceGetAPIRequest) GetNick() string {
    return r._nick
}
