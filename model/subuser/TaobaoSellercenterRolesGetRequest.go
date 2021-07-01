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
type TaobaoSellercenterRolesGetAPIRequest struct {
    model.Params
    // 卖家昵称(只允许查询自己的信息：当前登陆者)
    _nick   string
}

// 初始化TaobaoSellercenterRolesGetAPIRequest对象
func NewTaobaoSellercenterRolesGetRequest() *TaobaoSellercenterRolesGetAPIRequest{
    return &TaobaoSellercenterRolesGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSellercenterRolesGetAPIRequest) GetApiMethodName() string {
    return "taobao.sellercenter.roles.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSellercenterRolesGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 卖家昵称(只允许查询自己的信息：当前登陆者)
func (r *TaobaoSellercenterRolesGetAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSellercenterRolesGetAPIRequest) GetNick() string {
    return r._nick
}
