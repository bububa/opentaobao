package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclcancelpermiitemfromroleAPIRequest 取消角色和权限之间的关系 API请求
// alibaba.campus.acl.cancelpermiitemfromrole
//
// 取消角色和权限之间的关系
type AlibabacampusaclcancelpermiitemfromroleAPIRequest struct {
	model.Params
	// 系统自动生成
	_param2 []PermissionReq
	// 系统id
	_systemId string
	// 操作人id(不填默认appCode)
	_userId string
	// 公司ID
	_companyId int64
	// 园区id
	_campusId int64
	// 系统自动生成
	_param1 *RoleReq
}

// NewAlibabacampusaclcancelpermiitemfromroleRequest 初始化AlibabacampusaclcancelpermiitemfromroleAPIRequest对象
func NewAlibabacampusaclcancelpermiitemfromroleRequest() *AlibabacampusaclcancelpermiitemfromroleAPIRequest {
	return &AlibabacampusaclcancelpermiitemfromroleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusaclcancelpermiitemfromroleAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.cancelpermiitemfromrole"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusaclcancelpermiitemfromroleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusaclcancelpermiitemfromroleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam2 is Param2 Setter
// 系统自动生成
func (r *AlibabacampusaclcancelpermiitemfromroleAPIRequest) SetParam2(_param2 []PermissionReq) error {
	r._param2 = _param2
	r.Set("param2", _param2)
	return nil
}

// GetParam2 Param2 Getter
func (r AlibabacampusaclcancelpermiitemfromroleAPIRequest) GetParam2() []PermissionReq {
	return r._param2
}

// SetSystemId is SystemId Setter
// 系统id
func (r *AlibabacampusaclcancelpermiitemfromroleAPIRequest) SetSystemId(_systemId string) error {
	r._systemId = _systemId
	r.Set("system_id", _systemId)
	return nil
}

// GetSystemId SystemId Getter
func (r AlibabacampusaclcancelpermiitemfromroleAPIRequest) GetSystemId() string {
	return r._systemId
}

// SetUserId is UserId Setter
// 操作人id(不填默认appCode)
func (r *AlibabacampusaclcancelpermiitemfromroleAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabacampusaclcancelpermiitemfromroleAPIRequest) GetUserId() string {
	return r._userId
}

// SetCompanyId is CompanyId Setter
// 公司ID
func (r *AlibabacampusaclcancelpermiitemfromroleAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabacampusaclcancelpermiitemfromroleAPIRequest) GetCompanyId() int64 {
	return r._companyId
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabacampusaclcancelpermiitemfromroleAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabacampusaclcancelpermiitemfromroleAPIRequest) GetCampusId() int64 {
	return r._campusId
}

// SetParam1 is Param1 Setter
// 系统自动生成
func (r *AlibabacampusaclcancelpermiitemfromroleAPIRequest) SetParam1(_param1 *RoleReq) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabacampusaclcancelpermiitemfromroleAPIRequest) GetParam1() *RoleReq {
	return r._param1
}
