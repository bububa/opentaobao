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
type AlibabaMoziAclRoleAddRequest struct {
    model.Params
    // 创建角色请求对象
    _createRoleRequest   *CreateRoleRequest
}

// 初始化AlibabaMoziAclRoleAddRequest对象
func NewAlibabaMoziAclRoleAddRequest() *AlibabaMoziAclRoleAddRequest{
    return &AlibabaMoziAclRoleAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMoziAclRoleAddRequest) GetApiMethodName() string {
    return "alibaba.mozi.acl.role.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMoziAclRoleAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CreateRoleRequest Setter
// 创建角色请求对象
func (r *AlibabaMoziAclRoleAddRequest) SetCreateRoleRequest(_createRoleRequest *CreateRoleRequest) error {
    r._createRoleRequest = _createRoleRequest
    r.Set("create_role_request", _createRoleRequest)
    return nil
}

// CreateRoleRequest Getter
func (r AlibabaMoziAclRoleAddRequest) GetCreateRoleRequest() *CreateRoleRequest {
    return r._createRoleRequest
}
