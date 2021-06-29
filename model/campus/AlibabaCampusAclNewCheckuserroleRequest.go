package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
校验用户是否有角色 APIRequest
alibaba.campus.acl.new.checkuserrole

校验用户是否有角色
*/
type AlibabaCampusAclNewCheckuserroleRequest struct {
    model.Params

    // 用户账号
    userId   string 

    // 角色key
    roleKey   string 

    // 系统入参
    workbenchcontext   *WorkBenchContext 

}

func NewAlibabaCampusAclNewCheckuserroleRequest() *AlibabaCampusAclNewCheckuserroleRequest{
    return &AlibabaCampusAclNewCheckuserroleRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclNewCheckuserroleRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.checkuserrole"
}

func (r AlibabaCampusAclNewCheckuserroleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclNewCheckuserroleRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaCampusAclNewCheckuserroleRequest) GetUserId() string {
    return r.userId
}

func (r *AlibabaCampusAclNewCheckuserroleRequest) SetRoleKey(roleKey string) error {
    r.roleKey = roleKey
    r.Set("role_key", roleKey)
    return nil
}

func (r AlibabaCampusAclNewCheckuserroleRequest) GetRoleKey() string {
    return r.roleKey
}

func (r *AlibabaCampusAclNewCheckuserroleRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

func (r AlibabaCampusAclNewCheckuserroleRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}

