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
    _companyId   int64
    // 系统id
    _systemId   string
    // 园区id
    _campusId   int64
    // 用户id
    _empId   string
    // 权限
    _priv   []PermissionReq
    // 操作人id(不填默认appCode)
    _userId   string
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
func (r *AlibabaCampusAclGrantpermiitemstouserRequest) SetCompanyId(_companyId int64) error {
    r._companyId = _companyId
    r.Set("company_id", _companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaCampusAclGrantpermiitemstouserRequest) GetCompanyId() int64 {
    return r._companyId
}
// SystemId Setter
// 系统id
func (r *AlibabaCampusAclGrantpermiitemstouserRequest) SetSystemId(_systemId string) error {
    r._systemId = _systemId
    r.Set("system_id", _systemId)
    return nil
}

// SystemId Getter
func (r AlibabaCampusAclGrantpermiitemstouserRequest) GetSystemId() string {
    return r._systemId
}
// CampusId Setter
// 园区id
func (r *AlibabaCampusAclGrantpermiitemstouserRequest) SetCampusId(_campusId int64) error {
    r._campusId = _campusId
    r.Set("campus_id", _campusId)
    return nil
}

// CampusId Getter
func (r AlibabaCampusAclGrantpermiitemstouserRequest) GetCampusId() int64 {
    return r._campusId
}
// EmpId Setter
// 用户id
func (r *AlibabaCampusAclGrantpermiitemstouserRequest) SetEmpId(_empId string) error {
    r._empId = _empId
    r.Set("emp_id", _empId)
    return nil
}

// EmpId Getter
func (r AlibabaCampusAclGrantpermiitemstouserRequest) GetEmpId() string {
    return r._empId
}
// Priv Setter
// 权限
func (r *AlibabaCampusAclGrantpermiitemstouserRequest) SetPriv(_priv []PermissionReq) error {
    r._priv = _priv
    r.Set("priv", _priv)
    return nil
}

// Priv Getter
func (r AlibabaCampusAclGrantpermiitemstouserRequest) GetPriv() []PermissionReq {
    return r._priv
}
// UserId Setter
// 操作人id(不填默认appCode)
func (r *AlibabaCampusAclGrantpermiitemstouserRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaCampusAclGrantpermiitemstouserRequest) GetUserId() string {
    return r._userId
}
