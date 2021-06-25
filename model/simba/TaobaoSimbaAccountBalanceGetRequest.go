package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取实时余额，”元”为单位 APIRequest
taobao.simba.account.balance.get

获取实时余额，”元”为单位
*/
type TaobaoSimbaAccountBalanceGetRequest struct {
    model.Params

    // 主人昵称
    nick   string 

}

func NewTaobaoSimbaAccountBalanceGetRequest() *TaobaoSimbaAccountBalanceGetRequest{
    return &TaobaoSimbaAccountBalanceGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaAccountBalanceGetRequest) GetApiMethodName() string {
    return "taobao.simba.account.balance.get"
}

func (r TaobaoSimbaAccountBalanceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaAccountBalanceGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaAccountBalanceGetRequest) GetNick() string {
    return r.nick
}

