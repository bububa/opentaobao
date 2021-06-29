package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
解冻角色 APIRequest
alibaba.campus.acl.new.unfreezerole

解冻角色
*/
type AlibabaCampusAclNewUnfreezeroleRequest struct {
    model.Params

    // 系统参数
    workbenchcontext   *WorkBenchContext 

    // 角色主键id
    roleId   int64 

}

func NewAlibabaCampusAclNewUnfreezeroleRequest() *AlibabaCampusAclNewUnfreezeroleRequest{
    return &AlibabaCampusAclNewUnfreezeroleRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclNewUnfreezeroleRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.unfreezerole"
}

func (r AlibabaCampusAclNewUnfreezeroleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclNewUnfreezeroleRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

func (r AlibabaCampusAclNewUnfreezeroleRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}

func (r *AlibabaCampusAclNewUnfreezeroleRequest) SetRoleId(roleId int64) error {
    r.roleId = roleId
    r.Set("role_id", roleId)
    return nil
}

func (r AlibabaCampusAclNewUnfreezeroleRequest) GetRoleId() int64 {
    return r.roleId
}

