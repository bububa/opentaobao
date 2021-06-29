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
    _tribeId   int64
    // 用户信息
    _members   []OpenImUser
    // 用户信息
    _user   *OpenImUser
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
func (r *TaobaoOpenimTribeInviteRequest) SetTribeId(_tribeId int64) error {
    r._tribeId = _tribeId
    r.Set("tribe_id", _tribeId)
    return nil
}

// TribeId Getter
func (r TaobaoOpenimTribeInviteRequest) GetTribeId() int64 {
    return r._tribeId
}
// Members Setter
// 用户信息
func (r *TaobaoOpenimTribeInviteRequest) SetMembers(_members []OpenImUser) error {
    r._members = _members
    r.Set("members", _members)
    return nil
}

// Members Getter
func (r TaobaoOpenimTribeInviteRequest) GetMembers() []OpenImUser {
    return r._members
}
// User Setter
// 用户信息
func (r *TaobaoOpenimTribeInviteRequest) SetUser(_user *OpenImUser) error {
    r._user = _user
    r.Set("user", _user)
    return nil
}

// User Getter
func (r TaobaoOpenimTribeInviteRequest) GetUser() *OpenImUser {
    return r._user
}
