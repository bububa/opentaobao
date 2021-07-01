package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclGrantpermiitemtoroleAPIRequest
权限赋予角色 API请求
alibaba.campus.acl.grantpermiitemtorole

权限赋予角色 */
type AlibabaCampusAclGrantpermiitemtoroleAPIRequest struct {
	model.Params
	// 公司ID,不填统一初始化SYS_000
	_companyId int64
	// 系统id
	_systemId string
	// 园区ID
	_campusId int64
	// 系统自动生成
	_role *RoleReq
	// 系统自动生成
	_priv []PermissionReq
	// 操作人id
	_userId string
}

// NewAlibabaCampusAclGrantpermiitemtoroleRequest 初始化AlibabaCampusAclGrantpermiitemtoroleAPIRequest对象
func NewAlibabaCampusAclGrantpermiitemtoroleRequest() *AlibabaCampusAclGrantpermiitemtoroleAPIRequest {
	return &AlibabaCampusAclGrantpermiitemtoroleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclGrantpermiitemtoroleAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.grantpermiitemtorole"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclGrantpermiitemtoroleAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CompanyId Setter
// 公司ID,不填统一初始化SYS_000
func (r *AlibabaCampusAclGrantpermiitemtoroleAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// Get CompanyId Getter
func (r AlibabaCampusAclGrantpermiitemtoroleAPIRequest) GetCompanyId() int64 {
	return r._companyId
}

// Set is SystemId Setter
// 系统id
func (r *AlibabaCampusAclGrantpermiitemtoroleAPIRequest) SetSystemId(_systemId string) error {
	r._systemId = _systemId
	r.Set("system_id", _systemId)
	return nil
}

// Get SystemId Getter
func (r AlibabaCampusAclGrantpermiitemtoroleAPIRequest) GetSystemId() string {
	return r._systemId
}

// Set is CampusId Setter
// 园区ID
func (r *AlibabaCampusAclGrantpermiitemtoroleAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// Get CampusId Getter
func (r AlibabaCampusAclGrantpermiitemtoroleAPIRequest) GetCampusId() int64 {
	return r._campusId
}

// Set is Role Setter
// 系统自动生成
func (r *AlibabaCampusAclGrantpermiitemtoroleAPIRequest) SetRole(_role *RoleReq) error {
	r._role = _role
	r.Set("role", _role)
	return nil
}

// Get Role Getter
func (r AlibabaCampusAclGrantpermiitemtoroleAPIRequest) GetRole() *RoleReq {
	return r._role
}

// Set is Priv Setter
// 系统自动生成
func (r *AlibabaCampusAclGrantpermiitemtoroleAPIRequest) SetPriv(_priv []PermissionReq) error {
	r._priv = _priv
	r.Set("priv", _priv)
	return nil
}

// Get Priv Getter
func (r AlibabaCampusAclGrantpermiitemtoroleAPIRequest) GetPriv() []PermissionReq {
	return r._priv
}

// Set is UserId Setter
// 操作人id
func (r *AlibabaCampusAclGrantpermiitemtoroleAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r AlibabaCampusAclGrantpermiitemtoroleAPIRequest) GetUserId() string {
	return r._userId
}
