package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群设置管理员 APIRequest
taobao.openim.tribe.setmanager

OPENIM群设置管理员
*/
type TaobaoOpenimTribeSetmanagerRequest struct {
    model.Params

    // 用户信息
    user   *OpenImUser 

    // 群id
    tid   int64 

    // 用户信息
    member   *OpenImUser 

}

func NewTaobaoOpenimTribeSetmanagerRequest() *TaobaoOpenimTribeSetmanagerRequest{
    return &TaobaoOpenimTribeSetmanagerRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimTribeSetmanagerRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.setmanager"
}

func (r TaobaoOpenimTribeSetmanagerRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimTribeSetmanagerRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

func (r TaobaoOpenimTribeSetmanagerRequest) GetUser() *OpenImUser {
    return r.user
}

func (r *TaobaoOpenimTribeSetmanagerRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoOpenimTribeSetmanagerRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoOpenimTribeSetmanagerRequest) SetMember(member *OpenImUser) error {
    r.member = member
    r.Set("member", member)
    return nil
}

func (r TaobaoOpenimTribeSetmanagerRequest) GetMember() *OpenImUser {
    return r.member
}

