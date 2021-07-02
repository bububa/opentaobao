package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclGrantpermiitemstouserAPIRequest 给人直接授权 API请求
// alibaba.campus.acl.grantpermiitemstouser
//
// 给人直接授权
type AlibabaCampusAclGrantpermiitemstouserAPIRequest struct {
	model.Params
	// 公司id不填统一默认生成SYS_000
	_companyId int64
	// 系统id
	_systemId string
	// 园区id
	_campusId int64
	// 用户id
	_empId string
	// 权限
	_priv []PermissionReq
	// 操作人id(不填默认appCode)
	_userId string
}

// NewAlibabaCampusAclGrantpermiitemstouserRequest 初始化AlibabaCampusAclGrantpermiitemstouserAPIRequest对象
func NewAlibabaCampusAclGrantpermiitemstouserRequest() *AlibabaCampusAclGrantpermiitemstouserAPIRequest {
	return &AlibabaCampusAclGrantpermiitemstouserAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclGrantpermiitemstouserAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.grantpermiitemstouser"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclGrantpermiitemstouserAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCompanyId is CompanyId Setter
// 公司id不填统一默认生成SYS_000
func (r *AlibabaCampusAclGrantpermiitemstouserAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabaCampusAclGrantpermiitemstouserAPIRequest) GetCompanyId() int64 {
	return r._companyId
}

// SetSystemId is SystemId Setter
// 系统id
func (r *AlibabaCampusAclGrantpermiitemstouserAPIRequest) SetSystemId(_systemId string) error {
	r._systemId = _systemId
	r.Set("system_id", _systemId)
	return nil
}

// GetSystemId SystemId Getter
func (r AlibabaCampusAclGrantpermiitemstouserAPIRequest) GetSystemId() string {
	return r._systemId
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabaCampusAclGrantpermiitemstouserAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabaCampusAclGrantpermiitemstouserAPIRequest) GetCampusId() int64 {
	return r._campusId
}

// SetEmpId is EmpId Setter
// 用户id
func (r *AlibabaCampusAclGrantpermiitemstouserAPIRequest) SetEmpId(_empId string) error {
	r._empId = _empId
	r.Set("emp_id", _empId)
	return nil
}

// GetEmpId EmpId Getter
func (r AlibabaCampusAclGrantpermiitemstouserAPIRequest) GetEmpId() string {
	return r._empId
}

// SetPriv is Priv Setter
// 权限
func (r *AlibabaCampusAclGrantpermiitemstouserAPIRequest) SetPriv(_priv []PermissionReq) error {
	r._priv = _priv
	r.Set("priv", _priv)
	return nil
}

// GetPriv Priv Getter
func (r AlibabaCampusAclGrantpermiitemstouserAPIRequest) GetPriv() []PermissionReq {
	return r._priv
}

// SetUserId is UserId Setter
// 操作人id(不填默认appCode)
func (r *AlibabaCampusAclGrantpermiitemstouserAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaCampusAclGrantpermiitemstouserAPIRequest) GetUserId() string {
	return r._userId
}
