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
type TaobaoSimbaAccountBalanceGetRequest struct {
    model.Params
    // 主人昵称
    nick   string
}

// 初始化TaobaoSimbaAccountBalanceGetRequest对象
func NewTaobaoSimbaAccountBalanceGetRequest() *TaobaoSimbaAccountBalanceGetRequest{
    return &TaobaoSimbaAccountBalanceGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAccountBalanceGetRequest) GetApiMethodName() string {
    return "taobao.simba.account.balance.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAccountBalanceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaAccountBalanceGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaAccountBalanceGetRequest) GetNick() string {
    return r.nick
}
