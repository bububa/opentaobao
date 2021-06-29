package moziacl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
回收账户被授予的角色接口 API请求
alibaba.mozi.acl.userroles.revoke

调用此接口，会根据入参回收该账户下的该批角色
*/
type AlibabaMoziAclUserrolesRevokeRequest struct {
    model.Params
    // 回收角色入参对象
    revokeRolesRequest   *RevokeRolesRequest
}

// 初始化AlibabaMoziAclUserrolesRevokeRequest对象
func NewAlibabaMoziAclUserrolesRevokeRequest() *AlibabaMoziAclUserrolesRevokeRequest{
    return &AlibabaMoziAclUserrolesRevokeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMoziAclUserrolesRevokeRequest) GetApiMethodName() string {
    return "alibaba.mozi.acl.userroles.revoke"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMoziAclUserrolesRevokeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RevokeRolesRequest Setter
// 回收角色入参对象
func (r *AlibabaMoziAclUserrolesRevokeRequest) SetRevokeRolesRequest(revokeRolesRequest *RevokeRolesRequest) error {
    r.revokeRolesRequest = revokeRolesRequest
    r.Set("revoke_roles_request", revokeRolesRequest)
    return nil
}

// RevokeRolesRequest Getter
func (r AlibabaMoziAclUserrolesRevokeRequest) GetRevokeRolesRequest() *RevokeRolesRequest {
    return r.revokeRolesRequest
}
