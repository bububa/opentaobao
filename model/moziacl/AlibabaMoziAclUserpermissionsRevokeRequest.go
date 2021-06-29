package moziacl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
回收账户权限 APIRequest
alibaba.mozi.acl.userpermissions.revoke

调用此接口，会根据入参回收该账户下的该批权限点
*/
type AlibabaMoziAclUserpermissionsRevokeRequest struct {
    model.Params

    // 回收权限入参对象
    revokePermission   *RevokePermissionsRequest 

}

func NewAlibabaMoziAclUserpermissionsRevokeRequest() *AlibabaMoziAclUserpermissionsRevokeRequest{
    return &AlibabaMoziAclUserpermissionsRevokeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMoziAclUserpermissionsRevokeRequest) GetApiMethodName() string {
    return "alibaba.mozi.acl.userpermissions.revoke"
}

func (r AlibabaMoziAclUserpermissionsRevokeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMoziAclUserpermissionsRevokeRequest) SetRevokePermission(revokePermission *RevokePermissionsRequest) error {
    r.revokePermission = revokePermission
    r.Set("revoke_permission", revokePermission)
    return nil
}

func (r AlibabaMoziAclUserpermissionsRevokeRequest) GetRevokePermission() *RevokePermissionsRequest {
    return r.revokePermission
}

