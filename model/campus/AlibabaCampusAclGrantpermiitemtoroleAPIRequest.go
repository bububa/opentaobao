package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclGrantpermiitemtoroleAPIRequest 权限赋予角色 API请求
// alibaba.campus.acl.grantpermiitemtorole
//
// 权限赋予角色
type AlibabaCampusAclGrantpermiitemtoroleAPIRequest struct {
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

// NewAlibabaCampusAclGrantpermiitemtoroleRequest 初始化AlibabaCampusAclGrantpermiitemtoroleAPIRequest对象
func NewAlibabaCampusAclGrantpermiitemtoroleRequest() *AlibabaCampusAclGrantpermiitemtoroleAPIRequest {
	return &AlibabaCampusAclGrantpermiitemtoroleAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusAclGrantpermiitemtoroleAPIRequest) Reset() {
	r._priv = r._priv[:0]
	r._systemId = ""
	r._userId = ""
	r._companyId = 0
	r._campusId = 0
	r._role = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclGrantpermiitemtoroleAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.grantpermiitemtorole"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclGrantpermiitemtoroleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusAclGrantpermiitemtoroleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPriv is Priv Setter
// 系统自动生成
func (r *AlibabaCampusAclGrantpermiitemtoroleAPIRequest) SetPriv(_priv []PermissionReq) error {
	r._priv = _priv
	r.Set("priv", _priv)
	return nil
}

// GetPriv Priv Getter
func (r AlibabaCampusAclGrantpermiitemtoroleAPIRequest) GetPriv() []PermissionReq {
	return r._priv
}

// SetSystemId is SystemId Setter
// 系统id
func (r *AlibabaCampusAclGrantpermiitemtoroleAPIRequest) SetSystemId(_systemId string) error {
	r._systemId = _systemId
	r.Set("system_id", _systemId)
	return nil
}

// GetSystemId SystemId Getter
func (r AlibabaCampusAclGrantpermiitemtoroleAPIRequest) GetSystemId() string {
	return r._systemId
}

// SetUserId is UserId Setter
// 操作人id
func (r *AlibabaCampusAclGrantpermiitemtoroleAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaCampusAclGrantpermiitemtoroleAPIRequest) GetUserId() string {
	return r._userId
}

// SetCompanyId is CompanyId Setter
// 公司ID,不填统一初始化SYS_000
func (r *AlibabaCampusAclGrantpermiitemtoroleAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabaCampusAclGrantpermiitemtoroleAPIRequest) GetCompanyId() int64 {
	return r._companyId
}

// SetCampusId is CampusId Setter
// 园区ID
func (r *AlibabaCampusAclGrantpermiitemtoroleAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabaCampusAclGrantpermiitemtoroleAPIRequest) GetCampusId() int64 {
	return r._campusId
}

// SetRole is Role Setter
// 系统自动生成
func (r *AlibabaCampusAclGrantpermiitemtoroleAPIRequest) SetRole(_role *RoleReq) error {
	r._role = _role
	r.Set("role", _role)
	return nil
}

// GetRole Role Getter
func (r AlibabaCampusAclGrantpermiitemtoroleAPIRequest) GetRole() *RoleReq {
	return r._role
}

var poolAlibabaCampusAclGrantpermiitemtoroleAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusAclGrantpermiitemtoroleRequest()
	},
}

// GetAlibabaCampusAclGrantpermiitemtoroleRequest 从 sync.Pool 获取 AlibabaCampusAclGrantpermiitemtoroleAPIRequest
func GetAlibabaCampusAclGrantpermiitemtoroleAPIRequest() *AlibabaCampusAclGrantpermiitemtoroleAPIRequest {
	return poolAlibabaCampusAclGrantpermiitemtoroleAPIRequest.Get().(*AlibabaCampusAclGrantpermiitemtoroleAPIRequest)
}

// ReleaseAlibabaCampusAclGrantpermiitemtoroleAPIRequest 将 AlibabaCampusAclGrantpermiitemtoroleAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusAclGrantpermiitemtoroleAPIRequest(v *AlibabaCampusAclGrantpermiitemtoroleAPIRequest) {
	v.Reset()
	poolAlibabaCampusAclGrantpermiitemtoroleAPIRequest.Put(v)
}
