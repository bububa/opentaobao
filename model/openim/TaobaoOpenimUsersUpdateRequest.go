package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量更新用户信息 API请求
taobao.openim.users.update

批量更新用户信息
*/
type TaobaoOpenimUsersUpdateRequest struct {
    model.Params
    // 用户信息列表
    userinfos   []Userinfos
}

// 初始化TaobaoOpenimUsersUpdateRequest对象
func NewTaobaoOpenimUsersUpdateRequest() *TaobaoOpenimUsersUpdateRequest{
    return &TaobaoOpenimUsersUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimUsersUpdateRequest) GetApiMethodName() string {
    return "taobao.openim.users.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimUsersUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Userinfos Setter
// 用户信息列表
func (r *TaobaoOpenimUsersUpdateRequest) SetUserinfos(userinfos []Userinfos) error {
    r.userinfos = userinfos
    r.Set("userinfos", userinfos)
    return nil
}

// Userinfos Getter
func (r TaobaoOpenimUsersUpdateRequest) GetUserinfos() []Userinfos {
    return r.userinfos
}
