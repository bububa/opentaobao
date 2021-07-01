package moziacl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
将一个角色授予一个账号 API请求
alibaba.mozi.acl.grant.grantrole

根据入参，将入参中的角色授权给入参的某个账户，调用此接口后，该账户就会被授予该角色
*/
type AlibabaMoziAclGrantGrantroleAPIRequest struct {
    model.Params
    // 整体入参对象
    _grantRolesRequest   *GrantRolesRequest
}

// 初始化AlibabaMoziAclGrantGrantroleAPIRequest对象
func NewAlibabaMoziAclGrantGrantroleRequest() *AlibabaMoziAclGrantGrantroleAPIRequest{
    return &AlibabaMoziAclGrantGrantroleAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMoziAclGrantGrantroleAPIRequest) GetApiMethodName() string {
    return "alibaba.mozi.acl.grant.grantrole"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMoziAclGrantGrantroleAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GrantRolesRequest Setter
// 整体入参对象
func (r *AlibabaMoziAclGrantGrantroleAPIRequest) SetGrantRolesRequest(_grantRolesRequest *GrantRolesRequest) error {
    r._grantRolesRequest = _grantRolesRequest
    r.Set("grant_roles_request", _grantRolesRequest)
    return nil
}

// GrantRolesRequest Getter
func (r AlibabaMoziAclGrantGrantroleAPIRequest) GetGrantRolesRequest() *GrantRolesRequest {
    return r._grantRolesRequest
}
