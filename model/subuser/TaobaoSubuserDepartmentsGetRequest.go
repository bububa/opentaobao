package subuser

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定账户的所有部门列表 APIRequest
taobao.subuser.departments.get

获取指定账户的所有部门列表，其实包括有每个部门的ID、父部门ID、部门名称（通过主账号登陆只能查询属于该主账号下的所有部门信息）。
*/
type TaobaoSubuserDepartmentsGetRequest struct {
    model.Params

    // 主账号用户名
    userNick   string 

}

func NewTaobaoSubuserDepartmentsGetRequest() *TaobaoSubuserDepartmentsGetRequest{
    return &TaobaoSubuserDepartmentsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSubuserDepartmentsGetRequest) GetApiMethodName() string {
    return "taobao.subuser.departments.get"
}

func (r TaobaoSubuserDepartmentsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSubuserDepartmentsGetRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

func (r TaobaoSubuserDepartmentsGetRequest) GetUserNick() string {
    return r.userNick
}

