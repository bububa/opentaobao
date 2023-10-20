package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclinsertroleAPIRequest 新增角色 API请求
// alibaba.campus.acl.insertrole
//
// 新增角色
type AlibabacampusaclinsertroleAPIRequest struct {
	model.Params
	// 系统id
	_systemId string
	// 操作人id(不填默认appCode)
	_userId string
	// 角色描述
	_roleDesc string
	// 角色名称
	_roleName string
	// 默认人员角色.还有device设备角色类型
	_roleType string
	// 角色唯一ID,统一ROLE_开头,不填默认生成ROLE_UUID(32位随机数)
	_roleId string
	// 公司id,不填统一为SYS_000
	_companyId int64
	// 园区id
	_campusId int64
}

// NewAlibabacampusaclinsertroleRequest 初始化AlibabacampusaclinsertroleAPIRequest对象
func NewAlibabacampusaclinsertroleRequest() *AlibabacampusaclinsertroleAPIRequest {
	return &AlibabacampusaclinsertroleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusaclinsertroleAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.insertrole"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusaclinsertroleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusaclinsertroleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSystemId is SystemId Setter
// 系统id
func (r *AlibabacampusaclinsertroleAPIRequest) SetSystemId(_systemId string) error {
	r._systemId = _systemId
	r.Set("system_id", _systemId)
	return nil
}

// GetSystemId SystemId Getter
func (r AlibabacampusaclinsertroleAPIRequest) GetSystemId() string {
	return r._systemId
}

// SetUserId is UserId Setter
// 操作人id(不填默认appCode)
func (r *AlibabacampusaclinsertroleAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabacampusaclinsertroleAPIRequest) GetUserId() string {
	return r._userId
}

// SetRoleDesc is RoleDesc Setter
// 角色描述
func (r *AlibabacampusaclinsertroleAPIRequest) SetRoleDesc(_roleDesc string) error {
	r._roleDesc = _roleDesc
	r.Set("role_desc", _roleDesc)
	return nil
}

// GetRoleDesc RoleDesc Getter
func (r AlibabacampusaclinsertroleAPIRequest) GetRoleDesc() string {
	return r._roleDesc
}

// SetRoleName is RoleName Setter
// 角色名称
func (r *AlibabacampusaclinsertroleAPIRequest) SetRoleName(_roleName string) error {
	r._roleName = _roleName
	r.Set("role_name", _roleName)
	return nil
}

// GetRoleName RoleName Getter
func (r AlibabacampusaclinsertroleAPIRequest) GetRoleName() string {
	return r._roleName
}

// SetRoleType is RoleType Setter
// 默认人员角色.还有device设备角色类型
func (r *AlibabacampusaclinsertroleAPIRequest) SetRoleType(_roleType string) error {
	r._roleType = _roleType
	r.Set("role_type", _roleType)
	return nil
}

// GetRoleType RoleType Getter
func (r AlibabacampusaclinsertroleAPIRequest) GetRoleType() string {
	return r._roleType
}

// SetRoleId is RoleId Setter
// 角色唯一ID,统一ROLE_开头,不填默认生成ROLE_UUID(32位随机数)
func (r *AlibabacampusaclinsertroleAPIRequest) SetRoleId(_roleId string) error {
	r._roleId = _roleId
	r.Set("role_id", _roleId)
	return nil
}

// GetRoleId RoleId Getter
func (r AlibabacampusaclinsertroleAPIRequest) GetRoleId() string {
	return r._roleId
}

// SetCompanyId is CompanyId Setter
// 公司id,不填统一为SYS_000
func (r *AlibabacampusaclinsertroleAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabacampusaclinsertroleAPIRequest) GetCompanyId() int64 {
	return r._companyId
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabacampusaclinsertroleAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabacampusaclinsertroleAPIRequest) GetCampusId() int64 {
	return r._campusId
}
