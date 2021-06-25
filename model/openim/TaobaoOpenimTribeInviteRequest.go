package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群邀请加入 APIRequest
taobao.openim.tribe.invite

OPENIM群邀请加入接口
*/
type TaobaoOpenimTribeInviteRequest struct {
    model.Params

    // 群id
    tribeId   int64 

    // 用户信息
    members   []OpenImUser 

    // 用户信息
    user   *OpenImUser 

}

func NewTaobaoOpenimTribeInviteRequest() *TaobaoOpenimTribeInviteRequest{
    return &TaobaoOpenimTribeInviteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimTribeInviteRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.invite"
}

func (r TaobaoOpenimTribeInviteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimTribeInviteRequest) SetTribeId(tribeId int64) error {
    r.tribeId = tribeId
    r.Set("tribe_id", tribeId)
    return nil
}

func (r TaobaoOpenimTribeInviteRequest) GetTribeId() int64 {
    return r.tribeId
}

func (r *TaobaoOpenimTribeInviteRequest) SetMembers(members []OpenImUser) error {
    r.members = members
    r.Set("members", members)
    return nil
}

func (r TaobaoOpenimTribeInviteRequest) GetMembers() []OpenImUser {
    return r.members
}

func (r *TaobaoOpenimTribeInviteRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

func (r TaobaoOpenimTribeInviteRequest) GetUser() *OpenImUser {
    return r.user
}

