package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群成员获取 API请求
taobao.openim.tribe.getmembers

OPENIM群成员获取
*/
type TaobaoOpenimTribeGetmembersRequest struct {
    model.Params
    // 用户信息
    user   *OpenImUser
    // 群id
    tribeId   int64
}

// 初始化TaobaoOpenimTribeGetmembersRequest对象
func NewTaobaoOpenimTribeGetmembersRequest() *TaobaoOpenimTribeGetmembersRequest{
    return &TaobaoOpenimTribeGetmembersRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeGetmembersRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.getmembers"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeGetmembersRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// User Setter
// 用户信息
func (r *TaobaoOpenimTribeGetmembersRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

// User Getter
func (r TaobaoOpenimTribeGetmembersRequest) GetUser() *OpenImUser {
    return r.user
}
// TribeId Setter
// 群id
func (r *TaobaoOpenimTribeGetmembersRequest) SetTribeId(tribeId int64) error {
    r.tribeId = tribeId
    r.Set("tribe_id", tribeId)
    return nil
}

// TribeId Getter
func (r TaobaoOpenimTribeGetmembersRequest) GetTribeId() int64 {
    return r.tribeId
}
