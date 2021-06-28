package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除用户 APIRequest
taobao.openim.users.delete

批量删除用户
*/
type TaobaoOpenimUsersDeleteRequest struct {
    model.Params

    // 需要删除的用户列表，多个用户用半角逗号分隔，最多一次可以删除100个用户
    userids   []string 

}

func NewTaobaoOpenimUsersDeleteRequest() *TaobaoOpenimUsersDeleteRequest{
    return &TaobaoOpenimUsersDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimUsersDeleteRequest) GetApiMethodName() string {
    return "taobao.openim.users.delete"
}

func (r TaobaoOpenimUsersDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimUsersDeleteRequest) SetUserids(userids []string) error {
    r.userids = userids
    r.Set("userids", userids)
    return nil
}

func (r TaobaoOpenimUsersDeleteRequest) GetUserids() []string {
    return r.userids
}

