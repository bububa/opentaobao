package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群设置管理员 API请求
taobao.openim.tribe.setmanager

OPENIM群设置管理员
*/
type TaobaoOpenimTribeSetmanagerRequest struct {
    model.Params
    // 用户信息
    user   *OpenImUser
    // 群id
    tid   int64
    // 用户信息
    member   *OpenImUser
}

// 初始化TaobaoOpenimTribeSetmanagerRequest对象
func NewTaobaoOpenimTribeSetmanagerRequest() *TaobaoOpenimTribeSetmanagerRequest{
    return &TaobaoOpenimTribeSetmanagerRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeSetmanagerRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.setmanager"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeSetmanagerRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// User Setter
// 用户信息
func (r *TaobaoOpenimTribeSetmanagerRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

// User Getter
func (r TaobaoOpenimTribeSetmanagerRequest) GetUser() *OpenImUser {
    return r.user
}
// Tid Setter
// 群id
func (r *TaobaoOpenimTribeSetmanagerRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoOpenimTribeSetmanagerRequest) GetTid() int64 {
    return r.tid
}
// Member Setter
// 用户信息
func (r *TaobaoOpenimTribeSetmanagerRequest) SetMember(member *OpenImUser) error {
    r.member = member
    r.Set("member", member)
    return nil
}

// Member Getter
func (r TaobaoOpenimTribeSetmanagerRequest) GetMember() *OpenImUser {
    return r.member
}
