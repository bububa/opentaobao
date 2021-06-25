package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群取消管理员 APIRequest
taobao.openim.tribe.unsetmanager

OPENIM群取消管理员
*/
type TaobaoOpenimTribeUnsetmanagerRequest struct {
    model.Params

    // 用户信息
    user   *OpenImUser 

    // 群id
    tid   int64 

    // 用户信息
    member   *OpenImUser 

}

func NewTaobaoOpenimTribeUnsetmanagerRequest() *TaobaoOpenimTribeUnsetmanagerRequest{
    return &TaobaoOpenimTribeUnsetmanagerRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimTribeUnsetmanagerRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.unsetmanager"
}

func (r TaobaoOpenimTribeUnsetmanagerRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimTribeUnsetmanagerRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

func (r TaobaoOpenimTribeUnsetmanagerRequest) GetUser() *OpenImUser {
    return r.user
}

func (r *TaobaoOpenimTribeUnsetmanagerRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoOpenimTribeUnsetmanagerRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoOpenimTribeUnsetmanagerRequest) SetMember(member *OpenImUser) error {
    r.member = member
    r.Set("member", member)
    return nil
}

func (r TaobaoOpenimTribeUnsetmanagerRequest) GetMember() *OpenImUser {
    return r.member
}

