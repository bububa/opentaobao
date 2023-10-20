package campus

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusAclNewCheckusermenuAPIRequest) Reset() {
	r._workbenchcontext = nil
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewCheckusermenuAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.checkusermenu"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewCheckusermenuAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusAclNewCheckusermenuAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkbenchcontext is Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewCheckusermenuAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// GetWorkbenchcontext Workbenchcontext Getter
func (r AlibabaCampusAclNewCheckusermenuAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// SetParam is Param Setter
// 入参
func (r *AlibabaCampusAclNewCheckusermenuAPIRequest) SetParam(_param *CheckUserMenuParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaCampusAclNewCheckusermenuAPIRequest) GetParam() *CheckUserMenuParam {
	return r._param
}

var poolAlibabaCampusAclNewCheckusermenuAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusAclNewCheckusermenuRequest()
	},
}

// GetAlibabaCampusAclNewCheckusermenuRequest 从 sync.Pool 获取 AlibabaCampusAclNewCheckusermenuAPIRequest
func GetAlibabaCampusAclNewCheckusermenuAPIRequest() *AlibabaCampusAclNewCheckusermenuAPIRequest {
	return poolAlibabaCampusAclNewCheckusermenuAPIRequest.Get().(*AlibabaCampusAclNewCheckusermenuAPIRequest)
}

// ReleaseAlibabaCampusAclNewCheckusermenuAPIRequest 将 AlibabaCampusAclNewCheckusermenuAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusAclNewCheckusermenuAPIRequest(v *AlibabaCampusAclNewCheckusermenuAPIRequest) {
	v.Reset()
	poolAlibabaCampusAclNewCheckusermenuAPIRequest.Put(v)
}
