package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群踢出成员 API请求
taobao.openim.tribe.expel

OPENIM群踢出成员
*/
type TaobaoOpenimTribeExpelRequest struct {
    model.Params
    // 用户信息
    _user   *OpenImUser
    // 群id
    _tribeId   int64
    // 用户信息
    _member   *OpenImUser
}

// 初始化TaobaoOpenimTribeExpelRequest对象
func NewTaobaoOpenimTribeExpelRequest() *TaobaoOpenimTribeExpelRequest{
    return &TaobaoOpenimTribeExpelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeExpelRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.expel"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeExpelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// User Setter
// 用户信息
func (r *TaobaoOpenimTribeExpelRequest) SetUser(_user *OpenImUser) error {
    r._user = _user
    r.Set("user", _user)
    return nil
}

// User Getter
func (r TaobaoOpenimTribeExpelRequest) GetUser() *OpenImUser {
    return r._user
}
// TribeId Setter
// 群id
func (r *TaobaoOpenimTribeExpelRequest) SetTribeId(_tribeId int64) error {
    r._tribeId = _tribeId
    r.Set("tribe_id", _tribeId)
    return nil
}

// TribeId Getter
func (r TaobaoOpenimTribeExpelRequest) GetTribeId() int64 {
    return r._tribeId
}
// Member Setter
// 用户信息
func (r *TaobaoOpenimTribeExpelRequest) SetMember(_member *OpenImUser) error {
    r._member = _member
    r.Set("member", _member)
    return nil
}

// Member Getter
func (r TaobaoOpenimTribeExpelRequest) GetMember() *OpenImUser {
    return r._member
}
