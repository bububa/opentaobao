package campus

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusAclNewRemoveroleAPIRequest) Reset() {
	r._param0 = nil
	r._roleId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewRemoveroleAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.removerole"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewRemoveroleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusAclNewRemoveroleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 系统入参
func (r *AlibabaCampusAclNewRemoveroleAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaCampusAclNewRemoveroleAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetRoleId is RoleId Setter
// 角色主键id
func (r *AlibabaCampusAclNewRemoveroleAPIRequest) SetRoleId(_roleId int64) error {
	r._roleId = _roleId
	r.Set("role_id", _roleId)
	return nil
}

// GetRoleId RoleId Getter
func (r AlibabaCampusAclNewRemoveroleAPIRequest) GetRoleId() int64 {
	return r._roleId
}

var poolAlibabaCampusAclNewRemoveroleAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusAclNewRemoveroleRequest()
	},
}

// GetAlibabaCampusAclNewRemoveroleRequest 从 sync.Pool 获取 AlibabaCampusAclNewRemoveroleAPIRequest
func GetAlibabaCampusAclNewRemoveroleAPIRequest() *AlibabaCampusAclNewRemoveroleAPIRequest {
	return poolAlibabaCampusAclNewRemoveroleAPIRequest.Get().(*AlibabaCampusAclNewRemoveroleAPIRequest)
}

// ReleaseAlibabaCampusAclNewRemoveroleAPIRequest 将 AlibabaCampusAclNewRemoveroleAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusAclNewRemoveroleAPIRequest(v *AlibabaCampusAclNewRemoveroleAPIRequest) {
	v.Reset()
	poolAlibabaCampusAclNewRemoveroleAPIRequest.Put(v)
}
