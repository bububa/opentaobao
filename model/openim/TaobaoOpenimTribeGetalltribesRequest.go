package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户群列表 APIRequest
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

func NewTaobaoOpenimTribeGetalltribesRequest() *TaobaoOpenimTribeGetalltribesRequest{
    return &TaobaoOpenimTribeGetalltribesRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimTribeGetalltribesRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.getalltribes"
}

func (r TaobaoOpenimTribeGetalltribesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimTribeGetalltribesRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

func (r TaobaoOpenimTribeGetalltribesRequest) GetUser() *OpenImUser {
    return r.user
}

func (r *TaobaoOpenimTribeGetalltribesRequest) SetTribeTypes(tribeTypes []int64) error {
    r.tribeTypes = tribeTypes
    r.Set("tribe_types", tribeTypes)
    return nil
}

func (r TaobaoOpenimTribeGetalltribesRequest) GetTribeTypes() []int64 {
    return r.tribeTypes
}

