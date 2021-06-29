package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除角色 APIRequest
alibaba.campus.acl.new.removerole

删除角色
*/
type AlibabaCampusAclNewRemoveroleRequest struct {
    model.Params

    // 系统入参
    param0   *WorkBenchContext 

    // 角色主键id
    roleId   int64 

}

func NewAlibabaCampusAclNewRemoveroleRequest() *AlibabaCampusAclNewRemoveroleRequest{
    return &AlibabaCampusAclNewRemoveroleRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclNewRemoveroleRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.removerole"
}

func (r AlibabaCampusAclNewRemoveroleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclNewRemoveroleRequest) SetParam0(param0 *WorkBenchContext) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaCampusAclNewRemoveroleRequest) GetParam0() *WorkBenchContext {
    return r.param0
}

func (r *AlibabaCampusAclNewRemoveroleRequest) SetRoleId(roleId int64) error {
    r.roleId = roleId
    r.Set("role_id", roleId)
    return nil
}

func (r AlibabaCampusAclNewRemoveroleRequest) GetRoleId() int64 {
    return r.roleId
}

