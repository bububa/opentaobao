package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取用户信息 APIRequest
taobao.openim.users.get

批量获取用户信息
*/
type TaobaoOpenimUsersGetRequest struct {
    model.Params

    // 用户id序列
    userids   []String 

}

func NewTaobaoOpenimUsersGetRequest() *TaobaoOpenimUsersGetRequest{
    return &TaobaoOpenimUsersGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimUsersGetRequest) GetApiMethodName() string {
    return "taobao.openim.users.get"
}

func (r TaobaoOpenimUsersGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimUsersGetRequest) SetUserids(userids []String) error {
    r.userids = userids
    r.Set("userids", userids)
    return nil
}

func (r TaobaoOpenimUsersGetRequest) GetUserids() []String {
    return r.userids
}

