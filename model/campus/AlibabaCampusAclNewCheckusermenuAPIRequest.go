package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclnewcheckusermenuAPIRequest 校验用户是否有菜单权限 API请求
// alibaba.campus.acl.new.checkusermenu
//
// 校验用户是否有菜单权限
type AlibabacampusaclnewcheckusermenuAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 入参
	_param *CheckUserMenuParam
}

// NewAlibabacampusaclnewcheckusermenuRequest 初始化AlibabacampusaclnewcheckusermenuAPIRequest对象
func NewAlibabacampusaclnewcheckusermenuRequest() *AlibabacampusaclnewcheckusermenuAPIRequest {
	return &AlibabacampusaclnewcheckusermenuAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusaclnewcheckusermenuAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.checkusermenu"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusaclnewcheckusermenuAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusaclnewcheckusermenuAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkbenchcontext is Workbenchcontext Setter
// 系统入参
func (r *AlibabacampusaclnewcheckusermenuAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// GetWorkbenchcontext Workbenchcontext Getter
func (r AlibabacampusaclnewcheckusermenuAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// SetParam is Param Setter
// 入参
func (r *AlibabacampusaclnewcheckusermenuAPIRequest) SetParam(_param *CheckUserMenuParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabacampusaclnewcheckusermenuAPIRequest) GetParam() *CheckUserMenuParam {
	return r._param
}
