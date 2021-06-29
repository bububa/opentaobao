package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
冻结角色 APIRequest
alibaba.campus.acl.new.freezerole

冻结角色
*/
type AlibabaCampusAclNewFreezeroleRequest struct {
    model.Params

    // 系统入参
    workbenchcontext   *WorkBenchContext 

    // 角色主键id
    roleId   int64 

}

func NewAlibabaCampusAclNewFreezeroleRequest() *AlibabaCampusAclNewFreezeroleRequest{
    return &AlibabaCampusAclNewFreezeroleRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclNewFreezeroleRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.freezerole"
}

func (r AlibabaCampusAclNewFreezeroleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclNewFreezeroleRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

func (r AlibabaCampusAclNewFreezeroleRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}

func (r *AlibabaCampusAclNewFreezeroleRequest) SetRoleId(roleId int64) error {
    r.roleId = roleId
    r.Set("role_id", roleId)
    return nil
}

func (r AlibabaCampusAclNewFreezeroleRequest) GetRoleId() int64 {
    return r.roleId
}

