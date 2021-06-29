package moziacl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除角色 APIRequest
alibaba.mozi.acl.role.remove

根据传入的角色code、租户id，删除租户内对应的角色
*/
type AlibabaMoziAclRoleRemoveRequest struct {
    model.Params

    // 删除角色请求对象
    deleteRolesRequest   *DeleteRolesRequest 

}

func NewAlibabaMoziAclRoleRemoveRequest() *AlibabaMoziAclRoleRemoveRequest{
    return &AlibabaMoziAclRoleRemoveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMoziAclRoleRemoveRequest) GetApiMethodName() string {
    return "alibaba.mozi.acl.role.remove"
}

func (r AlibabaMoziAclRoleRemoveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMoziAclRoleRemoveRequest) SetDeleteRolesRequest(deleteRolesRequest *DeleteRolesRequest) error {
    r.deleteRolesRequest = deleteRolesRequest
    r.Set("delete_roles_request", deleteRolesRequest)
    return nil
}

func (r AlibabaMoziAclRoleRemoveRequest) GetDeleteRolesRequest() *DeleteRolesRequest {
    return r.deleteRolesRequest
}

