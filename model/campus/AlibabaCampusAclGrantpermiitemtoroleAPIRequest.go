package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclgrantpermiitemtoroleAPIRequest 权限赋予角色 API请求
// alibaba.campus.acl.grantpermiitemtorole
//
// 权限赋予角色
type AlibabacampusaclgrantpermiitemtoroleAPIRequest struct {
	model.Params
	// 系统自动生成
	_priv []PermissionReq
	// 系统id
	_systemId string
	// 操作人id
	_userId string
	// 公司ID,不填统一初始化SYS_000
	_companyId int64
	// 园区ID
	_campusId int64
	// 系统自动生成
	_role *RoleReq
}

// NewAlibabacampusaclgrantpermiitemtoroleRequest 初始化AlibabacampusaclgrantpermiitemtoroleAPIRequest对象
func NewAlibabacampusaclgrantpermiitemtoroleRequest() *AlibabacampusaclgrantpermiitemtoroleAPIRequest {
	return &AlibabacampusaclgrantpermiitemtoroleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusaclgrantpermiitemtoroleAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.grantpermiitemtorole"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusaclgrantpermiitemtoroleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusaclgrantpermiitemtoroleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPriv is Priv Setter
// 系统自动生成
func (r *AlibabacampusaclgrantpermiitemtoroleAPIRequest) SetPriv(_priv []PermissionReq) error {
	r._priv = _priv
	r.Set("priv", _priv)
	return nil
}

// GetPriv Priv Getter
func (r AlibabacampusaclgrantpermiitemtoroleAPIRequest) GetPriv() []PermissionReq {
	return r._priv
}

// SetSystemId is SystemId Setter
// 系统id
func (r *AlibabacampusaclgrantpermiitemtoroleAPIRequest) SetSystemId(_systemId string) error {
	r._systemId = _systemId
	r.Set("system_id", _systemId)
	return nil
}

// GetSystemId SystemId Getter
func (r AlibabacampusaclgrantpermiitemtoroleAPIRequest) GetSystemId() string {
	return r._systemId
}

// SetUserId is UserId Setter
// 操作人id
func (r *AlibabacampusaclgrantpermiitemtoroleAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabacampusaclgrantpermiitemtoroleAPIRequest) GetUserId() string {
	return r._userId
}

// SetCompanyId is CompanyId Setter
// 公司ID,不填统一初始化SYS_000
func (r *AlibabacampusaclgrantpermiitemtoroleAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabacampusaclgrantpermiitemtoroleAPIRequest) GetCompanyId() int64 {
	return r._companyId
}

// SetCampusId is CampusId Setter
// 园区ID
func (r *AlibabacampusaclgrantpermiitemtoroleAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabacampusaclgrantpermiitemtoroleAPIRequest) GetCampusId() int64 {
	return r._campusId
}

// SetRole is Role Setter
// 系统自动生成
func (r *AlibabacampusaclgrantpermiitemtoroleAPIRequest) SetRole(_role *RoleReq) error {
	r._role = _role
	r.Set("role", _role)
	return nil
}

// GetRole Role Getter
func (r AlibabacampusaclgrantpermiitemtoroleAPIRequest) GetRole() *RoleReq {
	return r._role
}
