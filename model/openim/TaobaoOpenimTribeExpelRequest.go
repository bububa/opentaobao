package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群踢出成员 APIRequest
taobao.openim.tribe.expel

OPENIM群踢出成员
*/
type TaobaoOpenimTribeExpelRequest struct {
    model.Params

    // 用户信息
    user   *OpenImUser 

    // 群id
    tribeId   int64 

    // 用户信息
    member   *OpenImUser 

}

func NewTaobaoOpenimTribeExpelRequest() *TaobaoOpenimTribeExpelRequest{
    return &TaobaoOpenimTribeExpelRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimTribeExpelRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.expel"
}

func (r TaobaoOpenimTribeExpelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimTribeExpelRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

func (r TaobaoOpenimTribeExpelRequest) GetUser() *OpenImUser {
    return r.user
}

func (r *TaobaoOpenimTribeExpelRequest) SetTribeId(tribeId int64) error {
    r.tribeId = tribeId
    r.Set("tribe_id", tribeId)
    return nil
}

func (r TaobaoOpenimTribeExpelRequest) GetTribeId() int64 {
    return r.tribeId
}

func (r *TaobaoOpenimTribeExpelRequest) SetMember(member *OpenImUser) error {
    r.member = member
    r.Set("member", member)
    return nil
}

func (r TaobaoOpenimTribeExpelRequest) GetMember() *OpenImUser {
    return r.member
}

