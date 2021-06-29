package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群邀请加入 API请求
taobao.openim.tribe.invite

OPENIM群邀请加入接口
*/
type TaobaoOpenimTribeInviteRequest struct {
    model.Params
    // 群id
    tribeId   int64
    // 用户信息
    members   []OpenImUser
    // 用户信息
    user   *OpenImUser
}

// 初始化TaobaoOpenimTribeInviteRequest对象
func NewTaobaoOpenimTribeInviteRequest() *TaobaoOpenimTribeInviteRequest{
    return &TaobaoOpenimTribeInviteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeInviteRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.invite"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeInviteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TribeId Setter
// 群id
func (r *TaobaoOpenimTribeInviteRequest) SetTribeId(tribeId int64) error {
    r.tribeId = tribeId
    r.Set("tribe_id", tribeId)
    return nil
}

// TribeId Getter
func (r TaobaoOpenimTribeInviteRequest) GetTribeId() int64 {
    return r.tribeId
}
// Members Setter
// 用户信息
func (r *TaobaoOpenimTribeInviteRequest) SetMembers(members []OpenImUser) error {
    r.members = members
    r.Set("members", members)
    return nil
}

// Members Getter
func (r TaobaoOpenimTribeInviteRequest) GetMembers() []OpenImUser {
    return r.members
}
// User Setter
// 用户信息
func (r *TaobaoOpenimTribeInviteRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

// User Getter
func (r TaobaoOpenimTribeInviteRequest) GetUser() *OpenImUser {
    return r.user
}
