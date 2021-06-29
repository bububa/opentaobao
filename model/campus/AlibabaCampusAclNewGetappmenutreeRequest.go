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
    _workbenchcontext   *WorkBenchContext
    // 是否关联查询出菜单下的权限
    _withpermission   bool
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
func (r *AlibabaCampusAclNewGetappmenutreeRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
    r._workbenchcontext = _workbenchcontext
    r.Set("workbenchcontext", _workbenchcontext)
    return nil
}

// Workbenchcontext Getter
func (r AlibabaCampusAclNewGetappmenutreeRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r._workbenchcontext
}
// Withpermission Setter
// 是否关联查询出菜单下的权限
func (r *AlibabaCampusAclNewGetappmenutreeRequest) SetWithpermission(_withpermission bool) error {
    r._withpermission = _withpermission
    r.Set("withpermission", _withpermission)
    return nil
}

// Withpermission Getter
func (r AlibabaCampusAclNewGetappmenutreeRequest) GetWithpermission() bool {
    return r._withpermission
}