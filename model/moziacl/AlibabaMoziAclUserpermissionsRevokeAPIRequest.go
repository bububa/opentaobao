package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamoziacluserpermissionsrevokeAPIRequest 回收账户权限 API请求
// alibaba.mozi.acl.userpermissions.revoke
//
// 调用此接口，会根据入参回收该账户下的该批权限点
type AlibabamoziacluserpermissionsrevokeAPIRequest struct {
	model.Params
	// 回收权限入参对象
	_revokePermission *RevokePermissionsRequest
}

// NewAlibabamoziacluserpermissionsrevokeRequest 初始化AlibabamoziacluserpermissionsrevokeAPIRequest对象
func NewAlibabamoziacluserpermissionsrevokeRequest() *AlibabamoziacluserpermissionsrevokeAPIRequest {
	return &AlibabamoziacluserpermissionsrevokeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamoziacluserpermissionsrevokeAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.acl.userpermissions.revoke"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamoziacluserpermissionsrevokeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamoziacluserpermissionsrevokeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRevokePermission is RevokePermission Setter
// 回收权限入参对象
func (r *AlibabamoziacluserpermissionsrevokeAPIRequest) SetRevokePermission(_revokePermission *RevokePermissionsRequest) error {
	r._revokePermission = _revokePermission
	r.Set("revoke_permission", _revokePermission)
	return nil
}

// GetRevokePermission RevokePermission Getter
func (r AlibabamoziacluserpermissionsrevokeAPIRequest) GetRevokePermission() *RevokePermissionsRequest {
	return r._revokePermission
}
