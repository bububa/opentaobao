package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclnewremoveroleAPIRequest 删除角色 API请求
// alibaba.campus.acl.new.removerole
//
// 删除角色
type AlibabacampusaclnewremoveroleAPIRequest struct {
	model.Params
	// 系统入参
	_param0 *WorkBenchContext
	// 角色主键id
	_roleId int64
}

// NewAlibabacampusaclnewremoveroleRequest 初始化AlibabacampusaclnewremoveroleAPIRequest对象
func NewAlibabacampusaclnewremoveroleRequest() *AlibabacampusaclnewremoveroleAPIRequest {
	return &AlibabacampusaclnewremoveroleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusaclnewremoveroleAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.removerole"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusaclnewremoveroleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusaclnewremoveroleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 系统入参
func (r *AlibabacampusaclnewremoveroleAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabacampusaclnewremoveroleAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetRoleId is RoleId Setter
// 角色主键id
func (r *AlibabacampusaclnewremoveroleAPIRequest) SetRoleId(_roleId int64) error {
	r._roleId = _roleId
	r.Set("role_id", _roleId)
	return nil
}

// GetRoleId RoleId Getter
func (r AlibabacampusaclnewremoveroleAPIRequest) GetRoleId() int64 {
	return r._roleId
}
