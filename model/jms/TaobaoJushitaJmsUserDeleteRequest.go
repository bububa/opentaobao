package jms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除ONS消息同步用户 API请求
taobao.jushita.jms.user.delete

删除ONS消息同步用户，删除后用户的消息将不会推送到聚石塔的ONS中
*/
type TaobaoJushitaJmsUserDeleteRequest struct {
    model.Params
    // 需要停止同步消息的用户nick
    userNick   string
}

// 初始化TaobaoJushitaJmsUserDeleteRequest对象
func NewTaobaoJushitaJmsUserDeleteRequest() *TaobaoJushitaJmsUserDeleteRequest{
    return &TaobaoJushitaJmsUserDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJushitaJmsUserDeleteRequest) GetApiMethodName() string {
    return "taobao.jushita.jms.user.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJushitaJmsUserDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserNick Setter
// 需要停止同步消息的用户nick
func (r *TaobaoJushitaJmsUserDeleteRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

// UserNick Getter
func (r TaobaoJushitaJmsUserDeleteRequest) GetUserNick() string {
    return r.userNick
}
