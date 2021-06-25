package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群主动加入 APIRequest
taobao.openim.tribe.join

OPENIM群主动加入
*/
type TaobaoOpenimTribeJoinRequest struct {
    model.Params

    // 用户信息
    user   *OpenImUser 

    // 群id
    tribeId   int64 

}

func NewTaobaoOpenimTribeJoinRequest() *TaobaoOpenimTribeJoinRequest{
    return &TaobaoOpenimTribeJoinRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimTribeJoinRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.join"
}

func (r TaobaoOpenimTribeJoinRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimTribeJoinRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

func (r TaobaoOpenimTribeJoinRequest) GetUser() *OpenImUser {
    return r.user
}

func (r *TaobaoOpenimTribeJoinRequest) SetTribeId(tribeId int64) error {
    r.tribeId = tribeId
    r.Set("tribe_id", tribeId)
    return nil
}

func (r TaobaoOpenimTribeJoinRequest) GetTribeId() int64 {
    return r.tribeId
}

