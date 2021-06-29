package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取用户信息 API请求
taobao.openim.users.get

批量获取用户信息
*/
type TaobaoOpenimUsersGetRequest struct {
    model.Params
    // 用户id序列
    _userids   []string
}

// 初始化TaobaoOpenimUsersGetRequest对象
func NewTaobaoOpenimUsersGetRequest() *TaobaoOpenimUsersGetRequest{
    return &TaobaoOpenimUsersGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimUsersGetRequest) GetApiMethodName() string {
    return "taobao.openim.users.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimUsersGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Userids Setter
// 用户id序列
func (r *TaobaoOpenimUsersGetRequest) SetUserids(_userids []string) error {
    r._userids = _userids
    r.Set("userids", _userids)
    return nil
}

// Userids Getter
func (r TaobaoOpenimUsersGetRequest) GetUserids() []string {
    return r._userids
}
