package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
冻结角色 API请求
alibaba.campus.acl.new.freezerole

冻结角色
*/
type AlibabaCampusAclNewFreezeroleRequest struct {
    model.Params
    // 系统入参
    workbenchcontext   *WorkBenchContext
    // 角色主键id
    roleId   int64
}

// 初始化AlibabaCampusAclNewFreezeroleRequest对象
func NewAlibabaCampusAclNewFreezeroleRequest() *AlibabaCampusAclNewFreezeroleRequest{
    return &AlibabaCampusAclNewFreezeroleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewFreezeroleRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.freezerole"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewFreezeroleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewFreezeroleRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

// Workbenchcontext Getter
func (r AlibabaCampusAclNewFreezeroleRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}
// RoleId Setter
// 角色主键id
func (r *AlibabaCampusAclNewFreezeroleRequest) SetRoleId(roleId int64) error {
    r.roleId = roleId
    r.Set("role_id", roleId)
    return nil
}

// RoleId Getter
func (r AlibabaCampusAclNewFreezeroleRequest) GetRoleId() int64 {
    return r.roleId
}
