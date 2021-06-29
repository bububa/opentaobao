package moziacl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
回收账户权限 API请求
alibaba.mozi.acl.userpermissions.revoke

调用此接口，会根据入参回收该账户下的该批权限点
*/
type AlibabaMoziAclUserpermissionsRevokeRequest struct {
    model.Params
    // 回收权限入参对象
    _revokePermission   *RevokePermissionsRequest
}

// 初始化AlibabaMoziAclUserpermissionsRevokeRequest对象
func NewAlibabaMoziAclUserpermissionsRevokeRequest() *AlibabaMoziAclUserpermissionsRevokeRequest{
    return &AlibabaMoziAclUserpermissionsRevokeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMoziAclUserpermissionsRevokeRequest) GetApiMethodName() string {
    return "alibaba.mozi.acl.userpermissions.revoke"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMoziAclUserpermissionsRevokeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RevokePermission Setter
// 回收权限入参对象
func (r *AlibabaMoziAclUserpermissionsRevokeRequest) SetRevokePermission(_revokePermission *RevokePermissionsRequest) error {
    r._revokePermission = _revokePermission
    r.Set("revoke_permission", _revokePermission)
    return nil
}

// RevokePermission Getter
func (r AlibabaMoziAclUserpermissionsRevokeRequest) GetRevokePermission() *RevokePermissionsRequest {
    return r._revokePermission
}
