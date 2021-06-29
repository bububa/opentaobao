package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
给人直接授权 API请求
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

// 初始化AlibabaCampusAclGrantpermiitemstouserRequest对象
func NewAlibabaCampusAclGrantpermiitemstouserRequest() *AlibabaCampusAclGrantpermiitemstouserRequest{
    return &AlibabaCampusAclGrantpermiitemstouserRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclGrantpermiitemstouserRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.grantpermiitemstouser"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclGrantpermiitemstouserRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CompanyId Setter
// 公司id不填统一默认生成SYS_000
func (r *AlibabaCampusAclGrantpermiitemstouserRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaCampusAclGrantpermiitemstouserRequest) GetCompanyId() int64 {
    return r.companyId
}
// SystemId Setter
// 系统id
func (r *AlibabaCampusAclGrantpermiitemstouserRequest) SetSystemId(systemId string) error {
    r.systemId = systemId
    r.Set("system_id", systemId)
    return nil
}

// SystemId Getter
func (r AlibabaCampusAclGrantpermiitemstouserRequest) GetSystemId() string {
    return r.systemId
}
// CampusId Setter
// 园区id
func (r *AlibabaCampusAclGrantpermiitemstouserRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

// CampusId Getter
func (r AlibabaCampusAclGrantpermiitemstouserRequest) GetCampusId() int64 {
    return r.campusId
}
// EmpId Setter
// 用户id
func (r *AlibabaCampusAclGrantpermiitemstouserRequest) SetEmpId(empId string) error {
    r.empId = empId
    r.Set("emp_id", empId)
    return nil
}

// EmpId Getter
func (r AlibabaCampusAclGrantpermiitemstouserRequest) GetEmpId() string {
    return r.empId
}
// Priv Setter
// 权限
func (r *AlibabaCampusAclGrantpermiitemstouserRequest) SetPriv(priv []PermissionReq) error {
    r.priv = priv
    r.Set("priv", priv)
    return nil
}

// Priv Getter
func (r AlibabaCampusAclGrantpermiitemstouserRequest) GetPriv() []PermissionReq {
    return r.priv
}
// UserId Setter
// 操作人id(不填默认appCode)
func (r *AlibabaCampusAclGrantpermiitemstouserRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaCampusAclGrantpermiitemstouserRequest) GetUserId() string {
    return r.userId
}
