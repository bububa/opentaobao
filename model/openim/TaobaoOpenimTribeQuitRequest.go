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
type TaobaoOpenimTribeQuitRequest struct {
    model.Params
    // 用户信息
    user   *OpenImUser
    // 群id
    tribeId   int64
}

// 初始化TaobaoOpenimTribeQuitRequest对象
func NewTaobaoOpenimTribeQuitRequest() *TaobaoOpenimTribeQuitRequest{
    return &TaobaoOpenimTribeQuitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeQuitRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.quit"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeQuitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// User Setter
// 用户信息
func (r *TaobaoOpenimTribeQuitRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

// User Getter
func (r TaobaoOpenimTribeQuitRequest) GetUser() *OpenImUser {
    return r.user
}
// TribeId Setter
// 群id
func (r *TaobaoOpenimTribeQuitRequest) SetTribeId(tribeId int64) error {
    r.tribeId = tribeId
    r.Set("tribe_id", tribeId)
    return nil
}

// TribeId Getter
func (r TaobaoOpenimTribeQuitRequest) GetTribeId() int64 {
    return r.tribeId
}
