package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询管理员 APIRequest
alibaba.campus.acl.new.pageuserrole

新增用户和角色的关系
*/
type AlibabaCampusAclNewPageuserroleRequest struct {
    model.Params

    // 系统入参
    workbenchcontext   *WorkBenchContext 

    // 入参
    usersRoleQueryParam   *UsersRoleQueryParam 

}

func NewAlibabaCampusAclNewPageuserroleRequest() *AlibabaCampusAclNewPageuserroleRequest{
    return &AlibabaCampusAclNewPageuserroleRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclNewPageuserroleRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.pageuserrole"
}

func (r AlibabaCampusAclNewPageuserroleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclNewPageuserroleRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

func (r AlibabaCampusAclNewPageuserroleRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}

func (r *AlibabaCampusAclNewPageuserroleRequest) SetUsersRoleQueryParam(usersRoleQueryParam *UsersRoleQueryParam) error {
    r.usersRoleQueryParam = usersRoleQueryParam
    r.Set("users_role_query_param", usersRoleQueryParam)
    return nil
}

func (r AlibabaCampusAclNewPageuserroleRequest) GetUsersRoleQueryParam() *UsersRoleQueryParam {
    return r.usersRoleQueryParam
}

