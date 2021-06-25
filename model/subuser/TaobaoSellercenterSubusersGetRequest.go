package subuser

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询指定账户的子账号列表 APIRequest
taobao.sellercenter.subusers.get

根据主账号nick查询该账号下所有的子账号列表，只能查询属于自己的账号信息 (主账号以及所属子账号)
*/
type TaobaoSellercenterSubusersGetRequest struct {
    model.Params

    // 表示卖家昵称
    nick   string 

}

func NewTaobaoSellercenterSubusersGetRequest() *TaobaoSellercenterSubusersGetRequest{
    return &TaobaoSellercenterSubusersGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSellercenterSubusersGetRequest) GetApiMethodName() string {
    return "taobao.sellercenter.subusers.get"
}

func (r TaobaoSellercenterSubusersGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSellercenterSubusersGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSellercenterSubusersGetRequest) GetNick() string {
    return r.nick
}

