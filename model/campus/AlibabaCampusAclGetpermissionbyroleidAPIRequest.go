package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclgetpermissionbyroleidAPIRequest 根据角色Id查询权限 API请求
// alibaba.campus.acl.getpermissionbyroleid
//
// 根据角色查询权限
type AlibabacampusaclgetpermissionbyroleidAPIRequest struct {
	model.Params
	// 系统id
	_systemId string
	// 角色id
	_roleId string
	// 园区id
	_campusId int64
	// 公司id
	_companyId int64
}

// NewAlibabacampusaclgetpermissionbyroleidRequest 初始化AlibabacampusaclgetpermissionbyroleidAPIRequest对象
func NewAlibabacampusaclgetpermissionbyroleidRequest() *AlibabacampusaclgetpermissionbyroleidAPIRequest {
	return &AlibabacampusaclgetpermissionbyroleidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusaclgetpermissionbyroleidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.getpermissionbyroleid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusaclgetpermissionbyroleidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusaclgetpermissionbyroleidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSystemId is SystemId Setter
// 系统id
func (r *AlibabacampusaclgetpermissionbyroleidAPIRequest) SetSystemId(_systemId string) error {
	r._systemId = _systemId
	r.Set("system_id", _systemId)
	return nil
}

// GetSystemId SystemId Getter
func (r AlibabacampusaclgetpermissionbyroleidAPIRequest) GetSystemId() string {
	return r._systemId
}

// SetRoleId is RoleId Setter
// 角色id
func (r *AlibabacampusaclgetpermissionbyroleidAPIRequest) SetRoleId(_roleId string) error {
	r._roleId = _roleId
	r.Set("role_id", _roleId)
	return nil
}

// GetRoleId RoleId Getter
func (r AlibabacampusaclgetpermissionbyroleidAPIRequest) GetRoleId() string {
	return r._roleId
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabacampusaclgetpermissionbyroleidAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabacampusaclgetpermissionbyroleidAPIRequest) GetCampusId() int64 {
	return r._campusId
}

// SetCompanyId is CompanyId Setter
// 公司id
func (r *AlibabacampusaclgetpermissionbyroleidAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabacampusaclgetpermissionbyroleidAPIRequest) GetCompanyId() int64 {
	return r._companyId
}
