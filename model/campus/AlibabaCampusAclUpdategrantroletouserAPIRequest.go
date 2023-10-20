package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclupdategrantroletouserAPIRequest 修改用户到角色关系 API请求
// alibaba.campus.acl.updategrantroletouser
//
// 修改用户到角色关系
type AlibabacampusaclupdategrantroletouserAPIRequest struct {
	model.Params
	// 角色
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

// NewAlibabacampusaclupdategrantroletouserRequest 初始化AlibabacampusaclupdategrantroletouserAPIRequest对象
func NewAlibabacampusaclupdategrantroletouserRequest() *AlibabacampusaclupdategrantroletouserAPIRequest {
	return &AlibabacampusaclupdategrantroletouserAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusaclupdategrantroletouserAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.updategrantroletouser"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusaclupdategrantroletouserAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusaclupdategrantroletouserAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRole is Role Setter
// 角色
func (r *AlibabacampusaclupdategrantroletouserAPIRequest) SetRole(_role []RoleReq) error {
	r._role = _role
	r.Set("role", _role)
	return nil
}

// GetRole Role Getter
func (r AlibabacampusaclupdategrantroletouserAPIRequest) GetRole() []RoleReq {
	return r._role
}

// SetSystemId is SystemId Setter
// 系统id
func (r *AlibabacampusaclupdategrantroletouserAPIRequest) SetSystemId(_systemId string) error {
	r._systemId = _systemId
	r.Set("system_id", _systemId)
	return nil
}

// GetSystemId SystemId Getter
func (r AlibabacampusaclupdategrantroletouserAPIRequest) GetSystemId() string {
	return r._systemId
}

// SetUserId is UserId Setter
// 操作人id(不填默认appCode)
func (r *AlibabacampusaclupdategrantroletouserAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabacampusaclupdategrantroletouserAPIRequest) GetUserId() string {
	return r._userId
}

// SetAccountId is AccountId Setter
// 用户账号
func (r *AlibabacampusaclupdategrantroletouserAPIRequest) SetAccountId(_accountId string) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r AlibabacampusaclupdategrantroletouserAPIRequest) GetAccountId() string {
	return r._accountId
}

// SetCompanyId is CompanyId Setter
// 公司id
func (r *AlibabacampusaclupdategrantroletouserAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabacampusaclupdategrantroletouserAPIRequest) GetCompanyId() int64 {
	return r._companyId
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabacampusaclupdategrantroletouserAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabacampusaclupdategrantroletouserAPIRequest) GetCampusId() int64 {
	return r._campusId
}
