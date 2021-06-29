package moziacl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
将一个角色授予一个账号 APIRequest
alibaba.mozi.acl.grant.grantrole

根据入参，将入参中的角色授权给入参的某个账户，调用此接口后，该账户就会被授予该角色
*/
type AlibabaMoziAclGrantGrantroleRequest struct {
    model.Params

    // 整体入参对象
    grantRolesRequest   *GrantRolesRequest 

}

func NewAlibabaMoziAclGrantGrantroleRequest() *AlibabaMoziAclGrantGrantroleRequest{
    return &AlibabaMoziAclGrantGrantroleRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMoziAclGrantGrantroleRequest) GetApiMethodName() string {
    return "alibaba.mozi.acl.grant.grantrole"
}

func (r AlibabaMoziAclGrantGrantroleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMoziAclGrantGrantroleRequest) SetGrantRolesRequest(grantRolesRequest *GrantRolesRequest) error {
    r.grantRolesRequest = grantRolesRequest
    r.Set("grant_roles_request", grantRolesRequest)
    return nil
}

func (r AlibabaMoziAclGrantGrantroleRequest) GetGrantRolesRequest() *GrantRolesRequest {
    return r.grantRolesRequest
}

