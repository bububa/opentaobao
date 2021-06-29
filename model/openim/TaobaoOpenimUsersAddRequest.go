package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加用户 API请求
taobao.openim.users.add

导入用户
*/
type TaobaoOpenimUsersAddRequest struct {
    model.Params
    // 用户信息列表
    _userinfos   []Userinfos
}

// 初始化TaobaoOpenimUsersAddRequest对象
func NewTaobaoOpenimUsersAddRequest() *TaobaoOpenimUsersAddRequest{
    return &TaobaoOpenimUsersAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimUsersAddRequest) GetApiMethodName() string {
    return "taobao.openim.users.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimUsersAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Userinfos Setter
// 用户信息列表
func (r *TaobaoOpenimUsersAddRequest) SetUserinfos(_userinfos []Userinfos) error {
    r._userinfos = _userinfos
    r.Set("userinfos", _userinfos)
    return nil
}

// Userinfos Getter
func (r TaobaoOpenimUsersAddRequest) GetUserinfos() []Userinfos {
    return r._userinfos
}
