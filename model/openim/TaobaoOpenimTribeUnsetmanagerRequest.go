package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群取消管理员 API请求
taobao.openim.tribe.unsetmanager

OPENIM群取消管理员
*/
type TaobaoOpenimTribeUnsetmanagerRequest struct {
    model.Params
    // 用户信息
    _user   *OpenImUser
    // 群id
    _tid   int64
    // 用户信息
    _member   *OpenImUser
}

// 初始化TaobaoOpenimTribeUnsetmanagerRequest对象
func NewTaobaoOpenimTribeUnsetmanagerRequest() *TaobaoOpenimTribeUnsetmanagerRequest{
    return &TaobaoOpenimTribeUnsetmanagerRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeUnsetmanagerRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.unsetmanager"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeUnsetmanagerRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// User Setter
// 用户信息
func (r *TaobaoOpenimTribeUnsetmanagerRequest) SetUser(_user *OpenImUser) error {
    r._user = _user
    r.Set("user", _user)
    return nil
}

// User Getter
func (r TaobaoOpenimTribeUnsetmanagerRequest) GetUser() *OpenImUser {
    return r._user
}
// Tid Setter
// 群id
func (r *TaobaoOpenimTribeUnsetmanagerRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoOpenimTribeUnsetmanagerRequest) GetTid() int64 {
    return r._tid
}
// Member Setter
// 用户信息
func (r *TaobaoOpenimTribeUnsetmanagerRequest) SetMember(_member *OpenImUser) error {
    r._member = _member
    r.Set("member", _member)
    return nil
}

// Member Getter
func (r TaobaoOpenimTribeUnsetmanagerRequest) GetMember() *OpenImUser {
    return r._member
}
