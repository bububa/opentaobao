package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改用户到角色关系 API请求
alibaba.campus.acl.updategrantroletouser

修改用户到角色关系
*/
type AlibabaCampusAclUpdategrantroletouserAPIRequest struct {
    model.Params
    // 公司id
    _companyId   int64
    // 系统id
    _systemId   string
    // 园区id
    _campusId   int64
    // 用户账号
    _accountId   string
    // 角色
    _role   []RoleReq
    // 操作人id(不填默认appCode)
    _userId   string
}

// 初始化AlibabaCampusAclUpdategrantroletouserAPIRequest对象
func NewAlibabaCampusAclUpdategrantroletouserRequest() *AlibabaCampusAclUpdategrantroletouserAPIRequest{
    return &AlibabaCampusAclUpdategrantroletouserAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclUpdategrantroletouserAPIRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.updategrantroletouser"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclUpdategrantroletouserAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CompanyId Setter
// 公司id
func (r *AlibabaCampusAclUpdategrantroletouserAPIRequest) SetCompanyId(_companyId int64) error {
    r._companyId = _companyId
    r.Set("company_id", _companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaCampusAclUpdategrantroletouserAPIRequest) GetCompanyId() int64 {
    return r._companyId
}
// SystemId Setter
// 系统id
func (r *AlibabaCampusAclUpdategrantroletouserAPIRequest) SetSystemId(_systemId string) error {
    r._systemId = _systemId
    r.Set("system_id", _systemId)
    return nil
}

// SystemId Getter
func (r AlibabaCampusAclUpdategrantroletouserAPIRequest) GetSystemId() string {
    return r._systemId
}
// CampusId Setter
// 园区id
func (r *AlibabaCampusAclUpdategrantroletouserAPIRequest) SetCampusId(_campusId int64) error {
    r._campusId = _campusId
    r.Set("campus_id", _campusId)
    return nil
}

// CampusId Getter
func (r AlibabaCampusAclUpdategrantroletouserAPIRequest) GetCampusId() int64 {
    return r._campusId
}
// AccountId Setter
// 用户账号
func (r *AlibabaCampusAclUpdategrantroletouserAPIRequest) SetAccountId(_accountId string) error {
    r._accountId = _accountId
    r.Set("account_id", _accountId)
    return nil
}

// AccountId Getter
func (r AlibabaCampusAclUpdategrantroletouserAPIRequest) GetAccountId() string {
    return r._accountId
}
// Role Setter
// 角色
func (r *AlibabaCampusAclUpdategrantroletouserAPIRequest) SetRole(_role []RoleReq) error {
    r._role = _role
    r.Set("role", _role)
    return nil
}

// Role Getter
func (r AlibabaCampusAclUpdategrantroletouserAPIRequest) GetRole() []RoleReq {
    return r._role
}
// UserId Setter
// 操作人id(不填默认appCode)
func (r *AlibabaCampusAclUpdategrantroletouserAPIRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaCampusAclUpdategrantroletouserAPIRequest) GetUserId() string {
    return r._userId
}
