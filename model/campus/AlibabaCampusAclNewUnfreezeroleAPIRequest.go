package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewUnfreezeroleAPIRequest 解冻角色 API请求
// alibaba.campus.acl.new.unfreezerole
//
// 解冻角色
type AlibabaCampusAclNewUnfreezeroleAPIRequest struct {
	model.Params
	// 系统参数
	_workbenchcontext *WorkBenchContext
	// 角色主键id
	_roleId int64
}

// NewAlibabaCampusAclNewUnfreezeroleRequest 初始化AlibabaCampusAclNewUnfreezeroleAPIRequest对象
func NewAlibabaCampusAclNewUnfreezeroleRequest() *AlibabaCampusAclNewUnfreezeroleAPIRequest {
	return &AlibabaCampusAclNewUnfreezeroleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewUnfreezeroleAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.unfreezerole"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewUnfreezeroleAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Workbenchcontext Setter
// 系统参数
func (r *AlibabaCampusAclNewUnfreezeroleAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// Get Workbenchcontext Getter
func (r AlibabaCampusAclNewUnfreezeroleAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// Set is RoleId Setter
// 角色主键id
func (r *AlibabaCampusAclNewUnfreezeroleAPIRequest) SetRoleId(_roleId int64) error {
	r._roleId = _roleId
	r.Set("role_id", _roleId)
	return nil
}

// Get RoleId Getter
func (r AlibabaCampusAclNewUnfreezeroleAPIRequest) GetRoleId() int64 {
	return r._roleId
}
