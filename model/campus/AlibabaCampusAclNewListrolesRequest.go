package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询全部角色 APIRequest
alibaba.campus.acl.new.listroles

查询全部角色
*/
type AlibabaCampusAclNewListrolesRequest struct {
    model.Params

    // 系统入参
    workbenchcontext   *WorkBenchContext 

    // 入参
    rolequeryparam   *RoleQueryParam 

}

func NewAlibabaCampusAclNewListrolesRequest() *AlibabaCampusAclNewListrolesRequest{
    return &AlibabaCampusAclNewListrolesRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclNewListrolesRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.listroles"
}

func (r AlibabaCampusAclNewListrolesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclNewListrolesRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

func (r AlibabaCampusAclNewListrolesRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}

func (r *AlibabaCampusAclNewListrolesRequest) SetRolequeryparam(rolequeryparam *RoleQueryParam) error {
    r.rolequeryparam = rolequeryparam
    r.Set("rolequeryparam", rolequeryparam)
    return nil
}

func (r AlibabaCampusAclNewListrolesRequest) GetRolequeryparam() *RoleQueryParam {
    return r.rolequeryparam
}

