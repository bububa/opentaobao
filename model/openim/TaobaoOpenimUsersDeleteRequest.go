package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除用户 API请求
taobao.openim.users.delete

批量删除用户
*/
type TaobaoOpenimUsersDeleteRequest struct {
    model.Params
    // 需要删除的用户列表，多个用户用半角逗号分隔，最多一次可以删除100个用户
    userids   []string
}

// 初始化TaobaoOpenimUsersDeleteRequest对象
func NewTaobaoOpenimUsersDeleteRequest() *TaobaoOpenimUsersDeleteRequest{
    return &TaobaoOpenimUsersDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimUsersDeleteRequest) GetApiMethodName() string {
    return "taobao.openim.users.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimUsersDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Userids Setter
// 需要删除的用户列表，多个用户用半角逗号分隔，最多一次可以删除100个用户
func (r *TaobaoOpenimUsersDeleteRequest) SetUserids(userids []string) error {
    r.userids = userids
    r.Set("userids", userids)
    return nil
}

// Userids Getter
func (r TaobaoOpenimUsersDeleteRequest) GetUserids() []string {
    return r.userids
}
