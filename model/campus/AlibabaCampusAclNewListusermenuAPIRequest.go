package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewListusermenuAPIRequest 查询用户有权限的菜单树 API请求
// alibaba.campus.acl.new.listusermenu
//
// 查询用户有权限的菜单树
type AlibabaCampusAclNewListusermenuAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 用户账号
	_userId string
}

// NewAlibabaCampusAclNewListusermenuRequest 初始化AlibabaCampusAclNewListusermenuAPIRequest对象
func NewAlibabaCampusAclNewListusermenuRequest() *AlibabaCampusAclNewListusermenuAPIRequest {
	return &AlibabaCampusAclNewListusermenuAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewListusermenuAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.listusermenu"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewListusermenuAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewListusermenuAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// Get Workbenchcontext Getter
func (r AlibabaCampusAclNewListusermenuAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// Set is UserId Setter
// 用户账号
func (r *AlibabaCampusAclNewListusermenuAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r AlibabaCampusAclNewListusermenuAPIRequest) GetUserId() string {
	return r._userId
}
