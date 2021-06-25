package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量更新用户信息 APIRequest
taobao.openim.users.update

批量更新用户信息
*/
type TaobaoOpenimUsersUpdateRequest struct {
    model.Params

    // 用户信息列表
    userinfos   []Userinfos 

}

func NewTaobaoOpenimUsersUpdateRequest() *TaobaoOpenimUsersUpdateRequest{
    return &TaobaoOpenimUsersUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimUsersUpdateRequest) GetApiMethodName() string {
    return "taobao.openim.users.update"
}

func (r TaobaoOpenimUsersUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimUsersUpdateRequest) SetUserinfos(userinfos []Userinfos) error {
    r.userinfos = userinfos
    r.Set("userinfos", userinfos)
    return nil
}

func (r TaobaoOpenimUsersUpdateRequest) GetUserinfos() []Userinfos {
    return r.userinfos
}

