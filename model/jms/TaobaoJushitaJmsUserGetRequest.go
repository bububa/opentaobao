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
type TaobaoJushitaJmsUserGetRequest struct {
    model.Params
    // 需要查询的用户名
    userNick   string
}

// 初始化TaobaoJushitaJmsUserGetRequest对象
func NewTaobaoJushitaJmsUserGetRequest() *TaobaoJushitaJmsUserGetRequest{
    return &TaobaoJushitaJmsUserGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJushitaJmsUserGetRequest) GetApiMethodName() string {
    return "taobao.jushita.jms.user.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJushitaJmsUserGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserNick Setter
// 需要查询的用户名
func (r *TaobaoJushitaJmsUserGetRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

// UserNick Getter
func (r TaobaoJushitaJmsUserGetRequest) GetUserNick() string {
    return r.userNick
}
