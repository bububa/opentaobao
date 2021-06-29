package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询管理员 API请求
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

// 初始化AlibabaCampusAclNewPageuserroleRequest对象
func NewAlibabaCampusAclNewPageuserroleRequest() *AlibabaCampusAclNewPageuserroleRequest{
    return &AlibabaCampusAclNewPageuserroleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewPageuserroleRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.pageuserrole"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewPageuserroleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewPageuserroleRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

// Workbenchcontext Getter
func (r AlibabaCampusAclNewPageuserroleRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}
// UsersRoleQueryParam Setter
// 入参
func (r *AlibabaCampusAclNewPageuserroleRequest) SetUsersRoleQueryParam(usersRoleQueryParam *UsersRoleQueryParam) error {
    r.usersRoleQueryParam = usersRoleQueryParam
    r.Set("users_role_query_param", usersRoleQueryParam)
    return nil
}

// UsersRoleQueryParam Getter
func (r AlibabaCampusAclNewPageuserroleRequest) GetUsersRoleQueryParam() *UsersRoleQueryParam {
    return r.usersRoleQueryParam
}
