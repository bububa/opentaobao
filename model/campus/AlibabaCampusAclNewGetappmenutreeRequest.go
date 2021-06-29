package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询应用下的菜单树 APIRequest
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

func NewAlibabaCampusAclNewGetappmenutreeRequest() *AlibabaCampusAclNewGetappmenutreeRequest{
    return &AlibabaCampusAclNewGetappmenutreeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclNewGetappmenutreeRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.getappmenutree"
}

func (r AlibabaCampusAclNewGetappmenutreeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclNewGetappmenutreeRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

func (r AlibabaCampusAclNewGetappmenutreeRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}

func (r *AlibabaCampusAclNewGetappmenutreeRequest) SetWithpermission(withpermission bool) error {
    r.withpermission = withpermission
    r.Set("withpermission", withpermission)
    return nil
}

func (r AlibabaCampusAclNewGetappmenutreeRequest) GetWithpermission() bool {
    return r.withpermission
}

