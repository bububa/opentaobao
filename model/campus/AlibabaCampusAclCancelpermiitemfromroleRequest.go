package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消角色和权限之间的关系 APIRequest
alibaba.campus.acl.cancelpermiitemfromrole

取消角色和权限之间的关系
*/
type AlibabaCampusAclCancelpermiitemfromroleRequest struct {
    model.Params

    // 公司ID
    companyId   int64 

    // 系统id
    systemId   string 

    // 园区id
    campusId   int64 

    // 系统自动生成
    param1   *RoleReq 

    // 系统自动生成
    param2   []PermissionReq 

    // 操作人id(不填默认appCode)
    userId   string 

}

func NewAlibabaCampusAclCancelpermiitemfromroleRequest() *AlibabaCampusAclCancelpermiitemfromroleRequest{
    return &AlibabaCampusAclCancelpermiitemfromroleRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclCancelpermiitemfromroleRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.cancelpermiitemfromrole"
}

func (r AlibabaCampusAclCancelpermiitemfromroleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclCancelpermiitemfromroleRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

func (r AlibabaCampusAclCancelpermiitemfromroleRequest) GetCompanyId() int64 {
    return r.companyId
}

func (r *AlibabaCampusAclCancelpermiitemfromroleRequest) SetSystemId(systemId string) error {
    r.systemId = systemId
    r.Set("system_id", systemId)
    return nil
}

func (r AlibabaCampusAclCancelpermiitemfromroleRequest) GetSystemId() string {
    return r.systemId
}

func (r *AlibabaCampusAclCancelpermiitemfromroleRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

func (r AlibabaCampusAclCancelpermiitemfromroleRequest) GetCampusId() int64 {
    return r.campusId
}

func (r *AlibabaCampusAclCancelpermiitemfromroleRequest) SetParam1(param1 *RoleReq) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaCampusAclCancelpermiitemfromroleRequest) GetParam1() *RoleReq {
    return r.param1
}

func (r *AlibabaCampusAclCancelpermiitemfromroleRequest) SetParam2(param2 []PermissionReq) error {
    r.param2 = param2
    r.Set("param2", param2)
    return nil
}

func (r AlibabaCampusAclCancelpermiitemfromroleRequest) GetParam2() []PermissionReq {
    return r.param2
}

func (r *AlibabaCampusAclCancelpermiitemfromroleRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaCampusAclCancelpermiitemfromroleRequest) GetUserId() string {
    return r.userId
}

