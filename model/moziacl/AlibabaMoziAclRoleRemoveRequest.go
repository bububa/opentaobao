package moziacl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除角色 API请求
alibaba.mozi.acl.role.remove

根据传入的角色code、租户id，删除租户内对应的角色
*/
type AlibabaMoziAclRoleRemoveRequest struct {
    model.Params
    // 删除角色请求对象
    _deleteRolesRequest   *DeleteRolesRequest
}

// 初始化AlibabaMoziAclRoleRemoveRequest对象
func NewAlibabaMoziAclRoleRemoveRequest() *AlibabaMoziAclRoleRemoveRequest{
    return &AlibabaMoziAclRoleRemoveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMoziAclRoleRemoveRequest) GetApiMethodName() string {
    return "alibaba.mozi.acl.role.remove"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMoziAclRoleRemoveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeleteRolesRequest Setter
// 删除角色请求对象
func (r *AlibabaMoziAclRoleRemoveRequest) SetDeleteRolesRequest(_deleteRolesRequest *DeleteRolesRequest) error {
    r._deleteRolesRequest = _deleteRolesRequest
    r.Set("delete_roles_request", _deleteRolesRequest)
    return nil
}

// DeleteRolesRequest Getter
func (r AlibabaMoziAclRoleRemoveRequest) GetDeleteRolesRequest() *DeleteRolesRequest {
    return r._deleteRolesRequest
}
