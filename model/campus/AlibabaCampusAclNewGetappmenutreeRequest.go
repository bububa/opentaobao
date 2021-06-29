package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询应用下的菜单树 API请求
alibaba.campus.acl.new.getappmenutree

查询应用下的菜单树
*/
type AlibabaCampusAclNewGetappmenutreeRequest struct {
    model.Params
    // 系统入参
    workbenchcontext   *WorkBenchContext
    // 是否关联查询出菜单下的权限
    withpermission   bool
}

// 初始化AlibabaCampusAclNewGetappmenutreeRequest对象
func NewAlibabaCampusAclNewGetappmenutreeRequest() *AlibabaCampusAclNewGetappmenutreeRequest{
    return &AlibabaCampusAclNewGetappmenutreeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewGetappmenutreeRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.getappmenutree"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewGetappmenutreeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewGetappmenutreeRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

// Workbenchcontext Getter
func (r AlibabaCampusAclNewGetappmenutreeRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}
// Withpermission Setter
// 是否关联查询出菜单下的权限
func (r *AlibabaCampusAclNewGetappmenutreeRequest) SetWithpermission(withpermission bool) error {
    r.withpermission = withpermission
    r.Set("withpermission", withpermission)
    return nil
}

// Withpermission Getter
func (r AlibabaCampusAclNewGetappmenutreeRequest) GetWithpermission() bool {
    return r.withpermission
}
