package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设置群成员昵称 API请求
taobao.openim.tribe.setmembernick

设置群成员昵称，存在如下两种场景
1 群主或管理员设置群成员昵称，该操作有权限控制。只针对普通群的群主和管理员开发此功能；讨论组群主不支持此设置操作
2 群成员设置自己的昵称，该功能对群所有成员开放
*/
type TaobaoOpenimTribeSetmembernickRequest struct {
    model.Params
    // 发起设置昵称的操作者，如果是设置其他成员的昵称，只有普通组的群主和管理员有权限
    user   *User
    // 群id
    tribeId   int64
    // 被设置昵称的群成员
    member   *User
    // 设置的昵称
    nick   string
}

// 初始化TaobaoOpenimTribeSetmembernickRequest对象
func NewTaobaoOpenimTribeSetmembernickRequest() *TaobaoOpenimTribeSetmembernickRequest{
    return &TaobaoOpenimTribeSetmembernickRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeSetmembernickRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.setmembernick"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeSetmembernickRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// User Setter
// 发起设置昵称的操作者，如果是设置其他成员的昵称，只有普通组的群主和管理员有权限
func (r *TaobaoOpenimTribeSetmembernickRequest) SetUser(user *User) error {
    r.user = user
    r.Set("user", user)
    return nil
}

// User Getter
func (r TaobaoOpenimTribeSetmembernickRequest) GetUser() *User {
    return r.user
}
// TribeId Setter
// 群id
func (r *TaobaoOpenimTribeSetmembernickRequest) SetTribeId(tribeId int64) error {
    r.tribeId = tribeId
    r.Set("tribe_id", tribeId)
    return nil
}

// TribeId Getter
func (r TaobaoOpenimTribeSetmembernickRequest) GetTribeId() int64 {
    return r.tribeId
}
// Member Setter
// 被设置昵称的群成员
func (r *TaobaoOpenimTribeSetmembernickRequest) SetMember(member *User) error {
    r.member = member
    r.Set("member", member)
    return nil
}

// Member Getter
func (r TaobaoOpenimTribeSetmembernickRequest) GetMember() *User {
    return r.member
}
// Nick Setter
// 设置的昵称
func (r *TaobaoOpenimTribeSetmembernickRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoOpenimTribeSetmembernickRequest) GetNick() string {
    return r.nick
}
