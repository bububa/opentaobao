package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclgrantpermiitemstouserAPIRequest 给人直接授权 API请求
// alibaba.campus.acl.grantpermiitemstouser
//
// 给人直接授权
type AlibabacampusaclgrantpermiitemstouserAPIRequest struct {
	model.Params
	// 权限
	_priv []PermissionReq
	// 系统id
	_systemId string
	// 操作人id(不填默认appCode)
	_userId string
	// 用户id
	_empId string
	// 公司id不填统一默认生成SYS_000
	_companyId int64
	// 园区id
	_campusId int64
}

// NewAlibabacampusaclgrantpermiitemstouserRequest 初始化AlibabacampusaclgrantpermiitemstouserAPIRequest对象
func NewAlibabacampusaclgrantpermiitemstouserRequest() *AlibabacampusaclgrantpermiitemstouserAPIRequest {
	return &AlibabacampusaclgrantpermiitemstouserAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusaclgrantpermiitemstouserAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.grantpermiitemstouser"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusaclgrantpermiitemstouserAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusaclgrantpermiitemstouserAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPriv is Priv Setter
// 权限
func (r *AlibabacampusaclgrantpermiitemstouserAPIRequest) SetPriv(_priv []PermissionReq) error {
	r._priv = _priv
	r.Set("priv", _priv)
	return nil
}

// GetPriv Priv Getter
func (r AlibabacampusaclgrantpermiitemstouserAPIRequest) GetPriv() []PermissionReq {
	return r._priv
}

// SetSystemId is SystemId Setter
// 系统id
func (r *AlibabacampusaclgrantpermiitemstouserAPIRequest) SetSystemId(_systemId string) error {
	r._systemId = _systemId
	r.Set("system_id", _systemId)
	return nil
}

// GetSystemId SystemId Getter
func (r AlibabacampusaclgrantpermiitemstouserAPIRequest) GetSystemId() string {
	return r._systemId
}

// SetUserId is UserId Setter
// 操作人id(不填默认appCode)
func (r *AlibabacampusaclgrantpermiitemstouserAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabacampusaclgrantpermiitemstouserAPIRequest) GetUserId() string {
	return r._userId
}

// SetEmpId is EmpId Setter
// 用户id
func (r *AlibabacampusaclgrantpermiitemstouserAPIRequest) SetEmpId(_empId string) error {
	r._empId = _empId
	r.Set("emp_id", _empId)
	return nil
}

// GetEmpId EmpId Getter
func (r AlibabacampusaclgrantpermiitemstouserAPIRequest) GetEmpId() string {
	return r._empId
}

// SetCompanyId is CompanyId Setter
// 公司id不填统一默认生成SYS_000
func (r *AlibabacampusaclgrantpermiitemstouserAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabacampusaclgrantpermiitemstouserAPIRequest) GetCompanyId() int64 {
	return r._companyId
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabacampusaclgrantpermiitemstouserAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabacampusaclgrantpermiitemstouserAPIRequest) GetCampusId() int64 {
	return r._campusId
}
