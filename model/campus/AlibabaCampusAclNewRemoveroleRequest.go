package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除角色 API请求
alibaba.campus.acl.new.removerole

删除角色
*/
type AlibabaCampusAclNewRemoveroleRequest struct {
    model.Params
    // 系统入参
    _param0   *WorkBenchContext
    // 角色主键id
    _roleId   int64
}

// 初始化AlibabaCampusAclNewRemoveroleRequest对象
func NewAlibabaCampusAclNewRemoveroleRequest() *AlibabaCampusAclNewRemoveroleRequest{
    return &AlibabaCampusAclNewRemoveroleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewRemoveroleRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.removerole"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewRemoveroleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 系统入参
func (r *AlibabaCampusAclNewRemoveroleRequest) SetParam0(_param0 *WorkBenchContext) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaCampusAclNewRemoveroleRequest) GetParam0() *WorkBenchContext {
    return r._param0
}
// RoleId Setter
// 角色主键id
func (r *AlibabaCampusAclNewRemoveroleRequest) SetRoleId(_roleId int64) error {
    r._roleId = _roleId
    r.Set("role_id", _roleId)
    return nil
}

// RoleId Getter
func (r AlibabaCampusAclNewRemoveroleRequest) GetRoleId() int64 {
    return r._roleId
}
