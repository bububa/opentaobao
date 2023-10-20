package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamoziaclroleaddAPIRequest 新增一个角色 API请求
// alibaba.mozi.acl.role.add
//
// 新增一个角色
type AlibabamoziaclroleaddAPIRequest struct {
	model.Params
	// 创建角色请求对象
	_createRoleRequest *CreateRoleRequest
}

// NewAlibabamoziaclroleaddRequest 初始化AlibabamoziaclroleaddAPIRequest对象
func NewAlibabamoziaclroleaddRequest() *AlibabamoziaclroleaddAPIRequest {
	return &AlibabamoziaclroleaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamoziaclroleaddAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.acl.role.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamoziaclroleaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamoziaclroleaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateRoleRequest is CreateRoleRequest Setter
// 创建角色请求对象
func (r *AlibabamoziaclroleaddAPIRequest) SetCreateRoleRequest(_createRoleRequest *CreateRoleRequest) error {
	r._createRoleRequest = _createRoleRequest
	r.Set("create_role_request", _createRoleRequest)
	return nil
}

// GetCreateRoleRequest CreateRoleRequest Getter
func (r AlibabamoziaclroleaddAPIRequest) GetCreateRoleRequest() *CreateRoleRequest {
	return r._createRoleRequest
}
