package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群主动加入 API请求
taobao.openim.tribe.join

OPENIM群主动加入
*/
type TaobaoOpenimTribeJoinRequest struct {
    model.Params
    // 用户信息
    _user   *OpenImUser
    // 群id
    _tribeId   int64
}

// 初始化TaobaoOpenimTribeJoinRequest对象
func NewTaobaoOpenimTribeJoinRequest() *TaobaoOpenimTribeJoinRequest{
    return &TaobaoOpenimTribeJoinRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeJoinRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.join"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeJoinRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// User Setter
// 用户信息
func (r *TaobaoOpenimTribeJoinRequest) SetUser(_user *OpenImUser) error {
    r._user = _user
    r.Set("user", _user)
    return nil
}

// User Getter
func (r TaobaoOpenimTribeJoinRequest) GetUser() *OpenImUser {
    return r._user
}
// TribeId Setter
// 群id
func (r *TaobaoOpenimTribeJoinRequest) SetTribeId(_tribeId int64) error {
    r._tribeId = _tribeId
    r.Set("tribe_id", _tribeId)
    return nil
}

// TribeId Getter
func (r TaobaoOpenimTribeJoinRequest) GetTribeId() int64 {
    return r._tribeId
}
