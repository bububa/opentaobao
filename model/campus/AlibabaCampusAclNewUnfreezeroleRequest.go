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
    _workbenchcontext   *WorkBenchContext
    // 角色主键id
    _roleId   int64
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
func (r *AlibabaCampusAclNewUnfreezeroleRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
    r._workbenchcontext = _workbenchcontext
    r.Set("workbenchcontext", _workbenchcontext)
    return nil
}

// Workbenchcontext Getter
func (r AlibabaCampusAclNewUnfreezeroleRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r._workbenchcontext
}
// RoleId Setter
// 角色主键id
func (r *AlibabaCampusAclNewUnfreezeroleRequest) SetRoleId(_roleId int64) error {
    r._roleId = _roleId
    r.Set("role_id", _roleId)
    return nil
}

// RoleId Getter
func (r AlibabaCampusAclNewUnfreezeroleRequest) GetRoleId() int64 {
    return r._roleId
}
