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
    _workbenchcontext   *WorkBenchContext
    // 角色主键id
    _roleId   int64
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
func (r *AlibabaCampusAclNewFreezeroleRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
    r._workbenchcontext = _workbenchcontext
    r.Set("workbenchcontext", _workbenchcontext)
    return nil
}

// Workbenchcontext Getter
func (r AlibabaCampusAclNewFreezeroleRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r._workbenchcontext
}
// RoleId Setter
// 角色主键id
func (r *AlibabaCampusAclNewFreezeroleRequest) SetRoleId(_roleId int64) error {
    r._roleId = _roleId
    r.Set("role_id", _roleId)
    return nil
}

// RoleId Getter
func (r AlibabaCampusAclNewFreezeroleRequest) GetRoleId() int64 {
    return r._roleId
}
