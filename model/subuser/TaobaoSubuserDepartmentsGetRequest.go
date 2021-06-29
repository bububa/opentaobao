package subuser

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定账户的所有部门列表 API请求
taobao.subuser.departments.get

获取指定账户的所有部门列表，其实包括有每个部门的ID、父部门ID、部门名称（通过主账号登陆只能查询属于该主账号下的所有部门信息）。
*/
type TaobaoSubuserDepartmentsGetRequest struct {
    model.Params
    // 主账号用户名
    _userNick   string
}

// 初始化TaobaoSubuserDepartmentsGetRequest对象
func NewTaobaoSubuserDepartmentsGetRequest() *TaobaoSubuserDepartmentsGetRequest{
    return &TaobaoSubuserDepartmentsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSubuserDepartmentsGetRequest) GetApiMethodName() string {
    return "taobao.subuser.departments.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSubuserDepartmentsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserNick Setter
// 主账号用户名
func (r *TaobaoSubuserDepartmentsGetRequest) SetUserNick(_userNick string) error {
    r._userNick = _userNick
    r.Set("user_nick", _userNick)
    return nil
}

// UserNick Getter
func (r TaobaoSubuserDepartmentsGetRequest) GetUserNick() string {
    return r._userNick
}
