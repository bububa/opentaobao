package moziacl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增一个角色 APIRequest
alibaba.mozi.acl.role.add

新增一个角色
*/
type AlibabaMoziAclRoleAddRequest struct {
    model.Params

    // 创建角色请求对象
    createRoleRequest   *CreateRoleRequest 

}

func NewAlibabaMoziAclRoleAddRequest() *AlibabaMoziAclRoleAddRequest{
    return &AlibabaMoziAclRoleAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMoziAclRoleAddRequest) GetApiMethodName() string {
    return "alibaba.mozi.acl.role.add"
}

func (r AlibabaMoziAclRoleAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMoziAclRoleAddRequest) SetCreateRoleRequest(createRoleRequest *CreateRoleRequest) error {
    r.createRoleRequest = createRoleRequest
    r.Set("create_role_request", createRoleRequest)
    return nil
}

func (r AlibabaMoziAclRoleAddRequest) GetCreateRoleRequest() *CreateRoleRequest {
    return r.createRoleRequest
}

