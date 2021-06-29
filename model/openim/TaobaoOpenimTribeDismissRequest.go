package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群解散 API请求
taobao.openim.tribe.dismiss

OPENIM群解散
*/
type TaobaoOpenimTribeDismissRequest struct {
    model.Params
    // 用户信息
    user   *OpenImUser
    // 群id
    tribeId   int64
}

// 初始化TaobaoOpenimTribeDismissRequest对象
func NewTaobaoOpenimTribeDismissRequest() *TaobaoOpenimTribeDismissRequest{
    return &TaobaoOpenimTribeDismissRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeDismissRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.dismiss"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeDismissRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// User Setter
// 用户信息
func (r *TaobaoOpenimTribeDismissRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

// User Getter
func (r TaobaoOpenimTribeDismissRequest) GetUser() *OpenImUser {
    return r.user
}
// TribeId Setter
// 群id
func (r *TaobaoOpenimTribeDismissRequest) SetTribeId(tribeId int64) error {
    r.tribeId = tribeId
    r.Set("tribe_id", tribeId)
    return nil
}

// TribeId Getter
func (r TaobaoOpenimTribeDismissRequest) GetTribeId() int64 {
    return r.tribeId
}
