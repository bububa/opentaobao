package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclCancelpermiitemfromroleAPIRequest 取消角色和权限之间的关系 API请求
// alibaba.campus.acl.cancelpermiitemfromrole
//
// 取消角色和权限之间的关系
type AlibabaCampusAclCancelpermiitemfromroleAPIRequest struct {
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

// NewAlibabaCampusAclCancelpermiitemfromroleRequest 初始化AlibabaCampusAclCancelpermiitemfromroleAPIRequest对象
func NewAlibabaCampusAclCancelpermiitemfromroleRequest() *AlibabaCampusAclCancelpermiitemfromroleAPIRequest {
	return &AlibabaCampusAclCancelpermiitemfromroleAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusAclCancelpermiitemfromroleAPIRequest) Reset() {
	r._param2 = r._param2[:0]
	r._systemId = ""
	r._userId = ""
	r._companyId = 0
	r._campusId = 0
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclCancelpermiitemfromroleAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.cancelpermiitemfromrole"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclCancelpermiitemfromroleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusAclCancelpermiitemfromroleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam2 is Param2 Setter
// 系统自动生成
func (r *AlibabaCampusAclCancelpermiitemfromroleAPIRequest) SetParam2(_param2 []PermissionReq) error {
	r._param2 = _param2
	r.Set("param2", _param2)
	return nil
}

// GetParam2 Param2 Getter
func (r AlibabaCampusAclCancelpermiitemfromroleAPIRequest) GetParam2() []PermissionReq {
	return r._param2
}

// SetSystemId is SystemId Setter
// 系统id
func (r *AlibabaCampusAclCancelpermiitemfromroleAPIRequest) SetSystemId(_systemId string) error {
	r._systemId = _systemId
	r.Set("system_id", _systemId)
	return nil
}

// GetSystemId SystemId Getter
func (r AlibabaCampusAclCancelpermiitemfromroleAPIRequest) GetSystemId() string {
	return r._systemId
}

// SetUserId is UserId Setter
// 操作人id(不填默认appCode)
func (r *AlibabaCampusAclCancelpermiitemfromroleAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaCampusAclCancelpermiitemfromroleAPIRequest) GetUserId() string {
	return r._userId
}

// SetCompanyId is CompanyId Setter
// 公司ID
func (r *AlibabaCampusAclCancelpermiitemfromroleAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabaCampusAclCancelpermiitemfromroleAPIRequest) GetCompanyId() int64 {
	return r._companyId
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabaCampusAclCancelpermiitemfromroleAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabaCampusAclCancelpermiitemfromroleAPIRequest) GetCampusId() int64 {
	return r._campusId
}

// SetParam1 is Param1 Setter
// 系统自动生成
func (r *AlibabaCampusAclCancelpermiitemfromroleAPIRequest) SetParam1(_param1 *RoleReq) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaCampusAclCancelpermiitemfromroleAPIRequest) GetParam1() *RoleReq {
	return r._param1
}

var poolAlibabaCampusAclCancelpermiitemfromroleAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusAclCancelpermiitemfromroleRequest()
	},
}

// GetAlibabaCampusAclCancelpermiitemfromroleRequest 从 sync.Pool 获取 AlibabaCampusAclCancelpermiitemfromroleAPIRequest
func GetAlibabaCampusAclCancelpermiitemfromroleAPIRequest() *AlibabaCampusAclCancelpermiitemfromroleAPIRequest {
	return poolAlibabaCampusAclCancelpermiitemfromroleAPIRequest.Get().(*AlibabaCampusAclCancelpermiitemfromroleAPIRequest)
}

// ReleaseAlibabaCampusAclCancelpermiitemfromroleAPIRequest 将 AlibabaCampusAclCancelpermiitemfromroleAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusAclCancelpermiitemfromroleAPIRequest(v *AlibabaCampusAclCancelpermiitemfromroleAPIRequest) {
	v.Reset()
	poolAlibabaCampusAclCancelpermiitemfromroleAPIRequest.Put(v)
}
