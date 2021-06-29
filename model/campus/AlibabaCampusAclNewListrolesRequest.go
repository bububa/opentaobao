package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询全部角色 API请求
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

// 初始化AlibabaCampusAclNewListrolesRequest对象
func NewAlibabaCampusAclNewListrolesRequest() *AlibabaCampusAclNewListrolesRequest{
    return &AlibabaCampusAclNewListrolesRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewListrolesRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.listroles"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewListrolesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewListrolesRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

// Workbenchcontext Getter
func (r AlibabaCampusAclNewListrolesRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}
// Rolequeryparam Setter
// 入参
func (r *AlibabaCampusAclNewListrolesRequest) SetRolequeryparam(rolequeryparam *RoleQueryParam) error {
    r.rolequeryparam = rolequeryparam
    r.Set("rolequeryparam", rolequeryparam)
    return nil
}

// Rolequeryparam Getter
func (r AlibabaCampusAclNewListrolesRequest) GetRolequeryparam() *RoleQueryParam {
    return r.rolequeryparam
}
