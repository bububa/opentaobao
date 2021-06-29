package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
解冻角色 API请求
alibaba.campus.acl.new.unfreezerole

解冻角色
*/
type AlibabaCampusAclNewUnfreezeroleRequest struct {
    model.Params
    // 系统参数
    workbenchcontext   *WorkBenchContext
    // 角色主键id
    roleId   int64
}

// 初始化AlibabaCampusAclNewUnfreezeroleRequest对象
func NewAlibabaCampusAclNewUnfreezeroleRequest() *AlibabaCampusAclNewUnfreezeroleRequest{
    return &AlibabaCampusAclNewUnfreezeroleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewUnfreezeroleRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.unfreezerole"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewUnfreezeroleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Workbenchcontext Setter
// 系统参数
func (r *AlibabaCampusAclNewUnfreezeroleRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

// Workbenchcontext Getter
func (r AlibabaCampusAclNewUnfreezeroleRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}
// RoleId Setter
// 角色主键id
func (r *AlibabaCampusAclNewUnfreezeroleRequest) SetRoleId(roleId int64) error {
    r.roleId = roleId
    r.Set("role_id", roleId)
    return nil
}

// RoleId Getter
func (r AlibabaCampusAclNewUnfreezeroleRequest) GetRoleId() int64 {
    return r.roleId
}
