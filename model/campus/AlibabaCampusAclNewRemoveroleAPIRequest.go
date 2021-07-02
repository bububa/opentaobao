package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewRemoveroleAPIRequest 删除角色 API请求
// alibaba.campus.acl.new.removerole
//
// 删除角色
type AlibabaCampusAclNewRemoveroleAPIRequest struct {
	model.Params
	// 系统入参
	_param0 *WorkBenchContext
	// 角色主键id
	_roleId int64
}

// NewAlibabaCampusAclNewRemoveroleRequest 初始化AlibabaCampusAclNewRemoveroleAPIRequest对象
func NewAlibabaCampusAclNewRemoveroleRequest() *AlibabaCampusAclNewRemoveroleAPIRequest {
	return &AlibabaCampusAclNewRemoveroleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewRemoveroleAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.removerole"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewRemoveroleAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 系统入参
func (r *AlibabaCampusAclNewRemoveroleAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlibabaCampusAclNewRemoveroleAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// Set is RoleId Setter
// 角色主键id
func (r *AlibabaCampusAclNewRemoveroleAPIRequest) SetRoleId(_roleId int64) error {
	r._roleId = _roleId
	r.Set("role_id", _roleId)
	return nil
}

// Get RoleId Getter
func (r AlibabaCampusAclNewRemoveroleAPIRequest) GetRoleId() int64 {
	return r._roleId
}
