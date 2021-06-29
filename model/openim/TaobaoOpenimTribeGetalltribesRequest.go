package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户群列表 API请求
taobao.openim.tribe.getalltribes

OPENIM群服务获取用户群列表
*/
type TaobaoOpenimTribeGetalltribesRequest struct {
    model.Params
    // 用户信息
    user   *OpenImUser
    // 群类型
    tribeTypes   []int64
}

// 初始化TaobaoOpenimTribeGetalltribesRequest对象
func NewTaobaoOpenimTribeGetalltribesRequest() *TaobaoOpenimTribeGetalltribesRequest{
    return &TaobaoOpenimTribeGetalltribesRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeGetalltribesRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.getalltribes"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeGetalltribesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// User Setter
// 用户信息
func (r *TaobaoOpenimTribeGetalltribesRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

// User Getter
func (r TaobaoOpenimTribeGetalltribesRequest) GetUser() *OpenImUser {
    return r.user
}
// TribeTypes Setter
// 群类型
func (r *TaobaoOpenimTribeGetalltribesRequest) SetTribeTypes(tribeTypes []int64) error {
    r.tribeTypes = tribeTypes
    r.Set("tribe_types", tribeTypes)
    return nil
}

// TribeTypes Getter
func (r TaobaoOpenimTribeGetalltribesRequest) GetTribeTypes() []int64 {
    return r.tribeTypes
}
