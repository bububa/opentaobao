package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
撤销用户授予的角色 API请求
alibaba.campus.acl.cancelrolesfromuser

撤销用户授予的角色
*/
type AlibabaCampusAclCancelrolesfromuserRequest struct {
    model.Params
    // 系统自动生成
    _role   []RoleReq
    // 公司id
    _companyId   int64
    // 系统id
    _systemId   string
    // 园区id
    _campusId   int64
    // 用户账号
    _accountId   string
    // 操作人id(不填默认appCode)
    _userId   string
}

// 初始化AlibabaCampusAclCancelrolesfromuserRequest对象
func NewAlibabaCampusAclCancelrolesfromuserRequest() *AlibabaCampusAclCancelrolesfromuserRequest{
    return &AlibabaCampusAclCancelrolesfromuserRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclCancelrolesfromuserRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.cancelrolesfromuser"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclCancelrolesfromuserRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Role Setter
// 系统自动生成
func (r *AlibabaCampusAclCancelrolesfromuserRequest) SetRole(_role []RoleReq) error {
    r._role = _role
    r.Set("role", _role)
    return nil
}

// Role Getter
func (r AlibabaCampusAclCancelrolesfromuserRequest) GetRole() []RoleReq {
    return r._role
}
// CompanyId Setter
// 公司id
func (r *AlibabaCampusAclCancelrolesfromuserRequest) SetCompanyId(_companyId int64) error {
    r._companyId = _companyId
    r.Set("company_id", _companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaCampusAclCancelrolesfromuserRequest) GetCompanyId() int64 {
    return r._companyId
}
// SystemId Setter
// 系统id
func (r *AlibabaCampusAclCancelrolesfromuserRequest) SetSystemId(_systemId string) error {
    r._systemId = _systemId
    r.Set("system_id", _systemId)
    return nil
}

// SystemId Getter
func (r AlibabaCampusAclCancelrolesfromuserRequest) GetSystemId() string {
    return r._systemId
}
// CampusId Setter
// 园区id
func (r *AlibabaCampusAclCancelrolesfromuserRequest) SetCampusId(_campusId int64) error {
    r._campusId = _campusId
    r.Set("campus_id", _campusId)
    return nil
}

// CampusId Getter
func (r AlibabaCampusAclCancelrolesfromuserRequest) GetCampusId() int64 {
    return r._campusId
}
// AccountId Setter
// 用户账号
func (r *AlibabaCampusAclCancelrolesfromuserRequest) SetAccountId(_accountId string) error {
    r._accountId = _accountId
    r.Set("account_id", _accountId)
    return nil
}

// AccountId Getter
func (r AlibabaCampusAclCancelrolesfromuserRequest) GetAccountId() string {
    return r._accountId
}
// UserId Setter
// 操作人id(不填默认appCode)
func (r *AlibabaCampusAclCancelrolesfromuserRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaCampusAclCancelrolesfromuserRequest) GetUserId() string {
    return r._userId
}
