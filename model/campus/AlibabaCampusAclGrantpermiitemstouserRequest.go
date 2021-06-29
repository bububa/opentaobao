package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
给人直接授权 APIRequest
alibaba.campus.acl.grantpermiitemstouser

给人直接授权
*/
type AlibabaCampusAclGrantpermiitemstouserRequest struct {
    model.Params

    // 公司id不填统一默认生成SYS_000
    companyId   int64 

    // 系统id
    systemId   string 

    // 园区id
    campusId   int64 

    // 用户id
    empId   string 

    // 权限
    priv   []PermissionReq 

    // 操作人id(不填默认appCode)
    userId   string 

}

func NewAlibabaCampusAclGrantpermiitemstouserRequest() *AlibabaCampusAclGrantpermiitemstouserRequest{
    return &AlibabaCampusAclGrantpermiitemstouserRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclGrantpermiitemstouserRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.grantpermiitemstouser"
}

func (r AlibabaCampusAclGrantpermiitemstouserRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclGrantpermiitemstouserRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

func (r AlibabaCampusAclGrantpermiitemstouserRequest) GetCompanyId() int64 {
    return r.companyId
}

func (r *AlibabaCampusAclGrantpermiitemstouserRequest) SetSystemId(systemId string) error {
    r.systemId = systemId
    r.Set("system_id", systemId)
    return nil
}

func (r AlibabaCampusAclGrantpermiitemstouserRequest) GetSystemId() string {
    return r.systemId
}

func (r *AlibabaCampusAclGrantpermiitemstouserRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

func (r AlibabaCampusAclGrantpermiitemstouserRequest) GetCampusId() int64 {
    return r.campusId
}

func (r *AlibabaCampusAclGrantpermiitemstouserRequest) SetEmpId(empId string) error {
    r.empId = empId
    r.Set("emp_id", empId)
    return nil
}

func (r AlibabaCampusAclGrantpermiitemstouserRequest) GetEmpId() string {
    return r.empId
}

func (r *AlibabaCampusAclGrantpermiitemstouserRequest) SetPriv(priv []PermissionReq) error {
    r.priv = priv
    r.Set("priv", priv)
    return nil
}

func (r AlibabaCampusAclGrantpermiitemstouserRequest) GetPriv() []PermissionReq {
    return r.priv
}

func (r *AlibabaCampusAclGrantpermiitemstouserRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaCampusAclGrantpermiitemstouserRequest) GetUserId() string {
    return r.userId
}

