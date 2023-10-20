package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamoziaclgrantgrantroleAPIRequest 将一个角色授予一个账号 API请求
// alibaba.mozi.acl.grant.grantrole
//
// 根据入参，将入参中的角色授权给入参的某个账户，调用此接口后，该账户就会被授予该角色
type AlibabamoziaclgrantgrantroleAPIRequest struct {
	model.Params
	// 整体入参对象
	_grantRolesRequest *GrantRolesRequest
}

// NewAlibabamoziaclgrantgrantroleRequest 初始化AlibabamoziaclgrantgrantroleAPIRequest对象
func NewAlibabamoziaclgrantgrantroleRequest() *AlibabamoziaclgrantgrantroleAPIRequest {
	return &AlibabamoziaclgrantgrantroleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamoziaclgrantgrantroleAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.acl.grant.grantrole"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamoziaclgrantgrantroleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamoziaclgrantgrantroleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGrantRolesRequest is GrantRolesRequest Setter
// 整体入参对象
func (r *AlibabamoziaclgrantgrantroleAPIRequest) SetGrantRolesRequest(_grantRolesRequest *GrantRolesRequest) error {
	r._grantRolesRequest = _grantRolesRequest
	r.Set("grant_roles_request", _grantRolesRequest)
	return nil
}

// GetGrantRolesRequest GrantRolesRequest Getter
func (r AlibabamoziaclgrantgrantroleAPIRequest) GetGrantRolesRequest() *GrantRolesRequest {
	return r._grantRolesRequest
}
