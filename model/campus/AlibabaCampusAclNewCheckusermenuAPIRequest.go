package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewCheckusermenuAPIRequest 校验用户是否有菜单权限 API请求
// alibaba.campus.acl.new.checkusermenu
//
// 校验用户是否有菜单权限
type AlibabaCampusAclNewCheckusermenuAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 入参
	_param *CheckUserMenuParam
}

// NewAlibabaCampusAclNewCheckusermenuRequest 初始化AlibabaCampusAclNewCheckusermenuAPIRequest对象
func NewAlibabaCampusAclNewCheckusermenuRequest() *AlibabaCampusAclNewCheckusermenuAPIRequest {
	return &AlibabaCampusAclNewCheckusermenuAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewCheckusermenuAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.checkusermenu"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewCheckusermenuAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewCheckusermenuAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// Get Workbenchcontext Getter
func (r AlibabaCampusAclNewCheckusermenuAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// Set is Param Setter
// 入参
func (r *AlibabaCampusAclNewCheckusermenuAPIRequest) SetParam(_param *CheckUserMenuParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlibabaCampusAclNewCheckusermenuAPIRequest) GetParam() *CheckUserMenuParam {
	return r._param
}
