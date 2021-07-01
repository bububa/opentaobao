package moziacl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增一个角色 API请求
alibaba.mozi.acl.role.add

新增一个角色
*/
type AlibabaMoziAclRoleAddAPIRequest struct {
    model.Params
    // 创建角色请求对象
    _createRoleRequest   *CreateRoleRequest
}

// 初始化AlibabaMoziAclRoleAddAPIRequest对象
func NewAlibabaMoziAclRoleAddRequest() *AlibabaMoziAclRoleAddAPIRequest{
    return &AlibabaMoziAclRoleAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMoziAclRoleAddAPIRequest) GetApiMethodName() string {
    return "alibaba.mozi.acl.role.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMoziAclRoleAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CreateRoleRequest Setter
// 创建角色请求对象
func (r *AlibabaMoziAclRoleAddAPIRequest) SetCreateRoleRequest(_createRoleRequest *CreateRoleRequest) error {
    r._createRoleRequest = _createRoleRequest
    r.Set("create_role_request", _createRoleRequest)
    return nil
}

// CreateRoleRequest Getter
func (r AlibabaMoziAclRoleAddAPIRequest) GetCreateRoleRequest() *CreateRoleRequest {
    return r._createRoleRequest
}
