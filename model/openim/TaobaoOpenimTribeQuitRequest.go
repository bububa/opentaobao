package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群成员退出 API请求
taobao.openim.tribe.quit

OPENIM群成员退出
*/
type TaobaoOpenimTribeQuitAPIRequest struct {
    model.Params
    // 用户信息
    _user   *OpenImUser
    // 群id
    _tribeId   int64
}

// 初始化TaobaoOpenimTribeQuitAPIRequest对象
func NewTaobaoOpenimTribeQuitRequest() *TaobaoOpenimTribeQuitAPIRequest{
    return &TaobaoOpenimTribeQuitAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeQuitAPIRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.quit"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeQuitAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// User Setter
// 用户信息
func (r *TaobaoOpenimTribeQuitAPIRequest) SetUser(_user *OpenImUser) error {
    r._user = _user
    r.Set("user", _user)
    return nil
}

// User Getter
func (r TaobaoOpenimTribeQuitAPIRequest) GetUser() *OpenImUser {
    return r._user
}
// TribeId Setter
// 群id
func (r *TaobaoOpenimTribeQuitAPIRequest) SetTribeId(_tribeId int64) error {
    r._tribeId = _tribeId
    r.Set("tribe_id", _tribeId)
    return nil
}

// TribeId Getter
func (r TaobaoOpenimTribeQuitAPIRequest) GetTribeId() int64 {
    return r._tribeId
}
