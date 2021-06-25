package subuser

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定卖家的角色列表 APIRequest
taobao.sellercenter.roles.get

获取指定卖家的角色列表，只能获取属于登陆者自己的信息。
*/
type TaobaoSellercenterRolesGetRequest struct {
    model.Params

    // 卖家昵称(只允许查询自己的信息：当前登陆者)
    nick   string 

}

func NewTaobaoSellercenterRolesGetRequest() *TaobaoSellercenterRolesGetRequest{
    return &TaobaoSellercenterRolesGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSellercenterRolesGetRequest) GetApiMethodName() string {
    return "taobao.sellercenter.roles.get"
}

func (r TaobaoSellercenterRolesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSellercenterRolesGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSellercenterRolesGetRequest) GetNick() string {
    return r.nick
}

