package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群成员获取 APIRequest
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

func NewTaobaoOpenimTribeGetmembersRequest() *TaobaoOpenimTribeGetmembersRequest{
    return &TaobaoOpenimTribeGetmembersRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimTribeGetmembersRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.getmembers"
}

func (r TaobaoOpenimTribeGetmembersRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimTribeGetmembersRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

func (r TaobaoOpenimTribeGetmembersRequest) GetUser() *OpenImUser {
    return r.user
}

func (r *TaobaoOpenimTribeGetmembersRequest) SetTribeId(tribeId int64) error {
    r.tribeId = tribeId
    r.Set("tribe_id", tribeId)
    return nil
}

func (r TaobaoOpenimTribeGetmembersRequest) GetTribeId() int64 {
    return r.tribeId
}

