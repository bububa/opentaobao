package subuser

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定卖家的角色列表 API请求
taobao.sellercenter.roles.get

获取指定卖家的角色列表，只能获取属于登陆者自己的信息。
*/
type TaobaoSellercenterRolesGetRequest struct {
    model.Params
    // 卖家昵称(只允许查询自己的信息：当前登陆者)
    nick   string
}

// 初始化TaobaoSellercenterRolesGetRequest对象
func NewTaobaoSellercenterRolesGetRequest() *TaobaoSellercenterRolesGetRequest{
    return &TaobaoSellercenterRolesGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSellercenterRolesGetRequest) GetApiMethodName() string {
    return "taobao.sellercenter.roles.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSellercenterRolesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 卖家昵称(只允许查询自己的信息：当前登陆者)
func (r *TaobaoSellercenterRolesGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSellercenterRolesGetRequest) GetNick() string {
    return r.nick
}
