package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除管理员 APIRequest
alibaba.campus.acl.new.deleteuserrole

删除管理员
*/
type AlibabaCampusAclNewDeleteuserroleRequest struct {
    model.Params

    // 系统入参
    workbenchcontext   *WorkBenchContext 

    // 用户账号
    userId   string 

    // 角色id
    roleIds   []int64 

}

func NewAlibabaCampusAclNewDeleteuserroleRequest() *AlibabaCampusAclNewDeleteuserroleRequest{
    return &AlibabaCampusAclNewDeleteuserroleRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclNewDeleteuserroleRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.deleteuserrole"
}

func (r AlibabaCampusAclNewDeleteuserroleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclNewDeleteuserroleRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

func (r AlibabaCampusAclNewDeleteuserroleRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}

func (r *AlibabaCampusAclNewDeleteuserroleRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaCampusAclNewDeleteuserroleRequest) GetUserId() string {
    return r.userId
}

func (r *AlibabaCampusAclNewDeleteuserroleRequest) SetRoleIds(roleIds []int64) error {
    r.roleIds = roleIds
    r.Set("role_ids", roleIds)
    return nil
}

func (r AlibabaCampusAclNewDeleteuserroleRequest) GetRoleIds() []int64 {
    return r.roleIds
}

