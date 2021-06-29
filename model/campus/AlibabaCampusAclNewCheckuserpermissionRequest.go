package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
校验用户是否有权限 APIRequest
alibaba.campus.acl.new.checkuserpermission

校验用户是否有权限
*/
type AlibabaCampusAclNewCheckuserpermissionRequest struct {
    model.Params

    // 系统入参
    workbenchcontext   *WorkBenchContext 

    // 接口入参
    checkUserPermissionParam   *CheckUserPermissionParam 

}

func NewAlibabaCampusAclNewCheckuserpermissionRequest() *AlibabaCampusAclNewCheckuserpermissionRequest{
    return &AlibabaCampusAclNewCheckuserpermissionRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclNewCheckuserpermissionRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.checkuserpermission"
}

func (r AlibabaCampusAclNewCheckuserpermissionRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclNewCheckuserpermissionRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

func (r AlibabaCampusAclNewCheckuserpermissionRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}

func (r *AlibabaCampusAclNewCheckuserpermissionRequest) SetCheckUserPermissionParam(checkUserPermissionParam *CheckUserPermissionParam) error {
    r.checkUserPermissionParam = checkUserPermissionParam
    r.Set("check_user_permission_param", checkUserPermissionParam)
    return nil
}

func (r AlibabaCampusAclNewCheckuserpermissionRequest) GetCheckUserPermissionParam() *CheckUserPermissionParam {
    return r.checkUserPermissionParam
}

