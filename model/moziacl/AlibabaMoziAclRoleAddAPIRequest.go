package moziacl

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziAclRoleAddAPIRequest 新增一个角色 API请求
// alibaba.mozi.acl.role.add
//
// 新增一个角色
type AlibabaMoziAclRoleAddAPIRequest struct {
	model.Params
	// 创建角色请求对象
	_createRoleRequest *CreateRoleRequest
}

// NewAlibabaMoziAclRoleAddRequest 初始化AlibabaMoziAclRoleAddAPIRequest对象
func NewAlibabaMoziAclRoleAddRequest() *AlibabaMoziAclRoleAddAPIRequest {
	return &AlibabaMoziAclRoleAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMoziAclRoleAddAPIRequest) Reset() {
	r._createRoleRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziAclRoleAddAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.acl.role.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziAclRoleAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMoziAclRoleAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateRoleRequest is CreateRoleRequest Setter
// 创建角色请求对象
func (r *AlibabaMoziAclRoleAddAPIRequest) SetCreateRoleRequest(_createRoleRequest *CreateRoleRequest) error {
	r._createRoleRequest = _createRoleRequest
	r.Set("create_role_request", _createRoleRequest)
	return nil
}

// GetCreateRoleRequest CreateRoleRequest Getter
func (r AlibabaMoziAclRoleAddAPIRequest) GetCreateRoleRequest() *CreateRoleRequest {
	return r._createRoleRequest
}

var poolAlibabaMoziAclRoleAddAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMoziAclRoleAddRequest()
	},
}

// GetAlibabaMoziAclRoleAddRequest 从 sync.Pool 获取 AlibabaMoziAclRoleAddAPIRequest
func GetAlibabaMoziAclRoleAddAPIRequest() *AlibabaMoziAclRoleAddAPIRequest {
	return poolAlibabaMoziAclRoleAddAPIRequest.Get().(*AlibabaMoziAclRoleAddAPIRequest)
}

// ReleaseAlibabaMoziAclRoleAddAPIRequest 将 AlibabaMoziAclRoleAddAPIRequest 放入 sync.Pool
func ReleaseAlibabaMoziAclRoleAddAPIRequest(v *AlibabaMoziAclRoleAddAPIRequest) {
	v.Reset()
	poolAlibabaMoziAclRoleAddAPIRequest.Put(v)
}
