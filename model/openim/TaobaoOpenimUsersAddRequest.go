package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加用户 APIRequest
taobao.openim.users.add

导入用户
*/
type TaobaoOpenimUsersAddRequest struct {
    model.Params

    // 用户信息列表
    userinfos   []Userinfos 

}

func NewTaobaoOpenimUsersAddRequest() *TaobaoOpenimUsersAddRequest{
    return &TaobaoOpenimUsersAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimUsersAddRequest) GetApiMethodName() string {
    return "taobao.openim.users.add"
}

func (r TaobaoOpenimUsersAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimUsersAddRequest) SetUserinfos(userinfos []Userinfos) error {
    r.userinfos = userinfos
    r.Set("userinfos", userinfos)
    return nil
}

func (r TaobaoOpenimUsersAddRequest) GetUserinfos() []Userinfos {
    return r.userinfos
}

