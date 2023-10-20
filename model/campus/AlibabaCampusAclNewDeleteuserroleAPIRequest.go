package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclnewdeleteuserroleAPIRequest 删除管理员 API请求
// alibaba.campus.acl.new.deleteuserrole
//
// 删除管理员
type AlibabacampusaclnewdeleteuserroleAPIRequest struct {
	model.Params
	// 角色id
	_roleIds []string
	// 用户账号
	_userId string
	// 系统入参
	_workbenchcontext *WorkBenchContext
}

// NewAlibabacampusaclnewdeleteuserroleRequest 初始化AlibabacampusaclnewdeleteuserroleAPIRequest对象
func NewAlibabacampusaclnewdeleteuserroleRequest() *AlibabacampusaclnewdeleteuserroleAPIRequest {
	return &AlibabacampusaclnewdeleteuserroleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusaclnewdeleteuserroleAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.deleteuserrole"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusaclnewdeleteuserroleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusaclnewdeleteuserroleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRoleIds is RoleIds Setter
// 角色id
func (r *AlibabacampusaclnewdeleteuserroleAPIRequest) SetRoleIds(_roleIds []string) error {
	r._roleIds = _roleIds
	r.Set("role_ids", _roleIds)
	return nil
}

// GetRoleIds RoleIds Getter
func (r AlibabacampusaclnewdeleteuserroleAPIRequest) GetRoleIds() []string {
	return r._roleIds
}

// SetUserId is UserId Setter
// 用户账号
func (r *AlibabacampusaclnewdeleteuserroleAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabacampusaclnewdeleteuserroleAPIRequest) GetUserId() string {
	return r._userId
}

// SetWorkbenchcontext is Workbenchcontext Setter
// 系统入参
func (r *AlibabacampusaclnewdeleteuserroleAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// GetWorkbenchcontext Workbenchcontext Getter
func (r AlibabacampusaclnewdeleteuserroleAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}
