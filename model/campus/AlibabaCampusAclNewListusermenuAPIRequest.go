package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclnewlistusermenuAPIRequest 查询用户有权限的菜单树 API请求
// alibaba.campus.acl.new.listusermenu
//
// 查询用户有权限的菜单树
type AlibabacampusaclnewlistusermenuAPIRequest struct {
	model.Params
	// 用户账号
	_userId string
	// 系统入参
	_workbenchcontext *WorkBenchContext
}

// NewAlibabacampusaclnewlistusermenuRequest 初始化AlibabacampusaclnewlistusermenuAPIRequest对象
func NewAlibabacampusaclnewlistusermenuRequest() *AlibabacampusaclnewlistusermenuAPIRequest {
	return &AlibabacampusaclnewlistusermenuAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusaclnewlistusermenuAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.listusermenu"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusaclnewlistusermenuAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusaclnewlistusermenuAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserId is UserId Setter
// 用户账号
func (r *AlibabacampusaclnewlistusermenuAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabacampusaclnewlistusermenuAPIRequest) GetUserId() string {
	return r._userId
}

// SetWorkbenchcontext is Workbenchcontext Setter
// 系统入参
func (r *AlibabacampusaclnewlistusermenuAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// GetWorkbenchcontext Workbenchcontext Getter
func (r AlibabacampusaclnewlistusermenuAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}
