package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclcancelrolesfromuserAPIRequest 撤销用户授予的角色 API请求
// alibaba.campus.acl.cancelrolesfromuser
//
// 撤销用户授予的角色
type AlibabacampusaclcancelrolesfromuserAPIRequest struct {
	model.Params
	// 系统自动生成
	_role []RoleReq
	// 系统id
	_systemId string
	// 操作人id(不填默认appCode)
	_userId string
	// 用户账号
	_accountId string
	// 公司id
	_companyId int64
	// 园区id
	_campusId int64
}

// NewAlibabacampusaclcancelrolesfromuserRequest 初始化AlibabacampusaclcancelrolesfromuserAPIRequest对象
func NewAlibabacampusaclcancelrolesfromuserRequest() *AlibabacampusaclcancelrolesfromuserAPIRequest {
	return &AlibabacampusaclcancelrolesfromuserAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusaclcancelrolesfromuserAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.cancelrolesfromuser"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusaclcancelrolesfromuserAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusaclcancelrolesfromuserAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRole is Role Setter
// 系统自动生成
func (r *AlibabacampusaclcancelrolesfromuserAPIRequest) SetRole(_role []RoleReq) error {
	r._role = _role
	r.Set("role", _role)
	return nil
}

// GetRole Role Getter
func (r AlibabacampusaclcancelrolesfromuserAPIRequest) GetRole() []RoleReq {
	return r._role
}

// SetSystemId is SystemId Setter
// 系统id
func (r *AlibabacampusaclcancelrolesfromuserAPIRequest) SetSystemId(_systemId string) error {
	r._systemId = _systemId
	r.Set("system_id", _systemId)
	return nil
}

// GetSystemId SystemId Getter
func (r AlibabacampusaclcancelrolesfromuserAPIRequest) GetSystemId() string {
	return r._systemId
}

// SetUserId is UserId Setter
// 操作人id(不填默认appCode)
func (r *AlibabacampusaclcancelrolesfromuserAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabacampusaclcancelrolesfromuserAPIRequest) GetUserId() string {
	return r._userId
}

// SetAccountId is AccountId Setter
// 用户账号
func (r *AlibabacampusaclcancelrolesfromuserAPIRequest) SetAccountId(_accountId string) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r AlibabacampusaclcancelrolesfromuserAPIRequest) GetAccountId() string {
	return r._accountId
}

// SetCompanyId is CompanyId Setter
// 公司id
func (r *AlibabacampusaclcancelrolesfromuserAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabacampusaclcancelrolesfromuserAPIRequest) GetCompanyId() int64 {
	return r._companyId
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabacampusaclcancelrolesfromuserAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabacampusaclcancelrolesfromuserAPIRequest) GetCampusId() int64 {
	return r._campusId
}
