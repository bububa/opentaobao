package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
校验用户是否有权限 API请求
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

// 初始化AlibabaCampusAclNewCheckuserpermissionRequest对象
func NewAlibabaCampusAclNewCheckuserpermissionRequest() *AlibabaCampusAclNewCheckuserpermissionRequest{
    return &AlibabaCampusAclNewCheckuserpermissionRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewCheckuserpermissionRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.checkuserpermission"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewCheckuserpermissionRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewCheckuserpermissionRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

// Workbenchcontext Getter
func (r AlibabaCampusAclNewCheckuserpermissionRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}
// CheckUserPermissionParam Setter
// 接口入参
func (r *AlibabaCampusAclNewCheckuserpermissionRequest) SetCheckUserPermissionParam(checkUserPermissionParam *CheckUserPermissionParam) error {
    r.checkUserPermissionParam = checkUserPermissionParam
    r.Set("check_user_permission_param", checkUserPermissionParam)
    return nil
}

// CheckUserPermissionParam Getter
func (r AlibabaCampusAclNewCheckuserpermissionRequest) GetCheckUserPermissionParam() *CheckUserPermissionParam {
    return r.checkUserPermissionParam
}
