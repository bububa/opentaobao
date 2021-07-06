package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclGetpermissionbyroleidAPIRequest 根据角色Id查询权限 API请求
// alibaba.campus.acl.getpermissionbyroleid
//
// 根据角色查询权限
type AlibabaCampusAclGetpermissionbyroleidAPIRequest struct {
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

// NewAlibabaCampusAclGetpermissionbyroleidRequest 初始化AlibabaCampusAclGetpermissionbyroleidAPIRequest对象
func NewAlibabaCampusAclGetpermissionbyroleidRequest() *AlibabaCampusAclGetpermissionbyroleidAPIRequest {
	return &AlibabaCampusAclGetpermissionbyroleidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclGetpermissionbyroleidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.getpermissionbyroleid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclGetpermissionbyroleidAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSystemId is SystemId Setter
// 系统id
func (r *AlibabaCampusAclGetpermissionbyroleidAPIRequest) SetSystemId(_systemId string) error {
	r._systemId = _systemId
	r.Set("system_id", _systemId)
	return nil
}

// GetSystemId SystemId Getter
func (r AlibabaCampusAclGetpermissionbyroleidAPIRequest) GetSystemId() string {
	return r._systemId
}

// SetRoleId is RoleId Setter
// 角色id
func (r *AlibabaCampusAclGetpermissionbyroleidAPIRequest) SetRoleId(_roleId string) error {
	r._roleId = _roleId
	r.Set("role_id", _roleId)
	return nil
}

// GetRoleId RoleId Getter
func (r AlibabaCampusAclGetpermissionbyroleidAPIRequest) GetRoleId() string {
	return r._roleId
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabaCampusAclGetpermissionbyroleidAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabaCampusAclGetpermissionbyroleidAPIRequest) GetCampusId() int64 {
	return r._campusId
}

// SetCompanyId is CompanyId Setter
// 公司id
func (r *AlibabaCampusAclGetpermissionbyroleidAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabaCampusAclGetpermissionbyroleidAPIRequest) GetCompanyId() int64 {
	return r._companyId
}
