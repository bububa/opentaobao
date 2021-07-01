package jms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询某个用户是否同步消息 API请求
taobao.jushita.jms.user.get

查询某个用户是否同步消息，只支持单个查询
*/
type TaobaoJushitaJmsUserGetAPIRequest struct {
    model.Params
    // 需要查询的用户名
    _userNick   string
}

// 初始化TaobaoJushitaJmsUserGetAPIRequest对象
func NewTaobaoJushitaJmsUserGetRequest() *TaobaoJushitaJmsUserGetAPIRequest{
    return &TaobaoJushitaJmsUserGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJushitaJmsUserGetAPIRequest) GetApiMethodName() string {
    return "taobao.jushita.jms.user.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJushitaJmsUserGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserNick Setter
// 需要查询的用户名
func (r *TaobaoJushitaJmsUserGetAPIRequest) SetUserNick(_userNick string) error {
    r._userNick = _userNick
    r.Set("user_nick", _userNick)
    return nil
}

// UserNick Getter
func (r TaobaoJushitaJmsUserGetAPIRequest) GetUserNick() string {
    return r._userNick
}
