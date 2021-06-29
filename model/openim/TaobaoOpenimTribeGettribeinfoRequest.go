package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取群信息 API请求
taobao.openim.tribe.gettribeinfo

获取群信息
*/
type TaobaoOpenimTribeGettribeinfoRequest struct {
    model.Params
    // 用户信息
    user   *OpenImUser
    // 群ID
    tribeId   int64
}

// 初始化TaobaoOpenimTribeGettribeinfoRequest对象
func NewTaobaoOpenimTribeGettribeinfoRequest() *TaobaoOpenimTribeGettribeinfoRequest{
    return &TaobaoOpenimTribeGettribeinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeGettribeinfoRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.gettribeinfo"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeGettribeinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// User Setter
// 用户信息
func (r *TaobaoOpenimTribeGettribeinfoRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

// User Getter
func (r TaobaoOpenimTribeGettribeinfoRequest) GetUser() *OpenImUser {
    return r.user
}
// TribeId Setter
// 群ID
func (r *TaobaoOpenimTribeGettribeinfoRequest) SetTribeId(tribeId int64) error {
    r.tribeId = tribeId
    r.Set("tribe_id", tribeId)
    return nil
}

// TribeId Getter
func (r TaobaoOpenimTribeGettribeinfoRequest) GetTribeId() int64 {
    return r.tribeId
}
