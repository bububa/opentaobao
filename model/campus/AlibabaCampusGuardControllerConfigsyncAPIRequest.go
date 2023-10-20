package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusGuardControllerConfigsyncAPIRequest 门禁控制器配置项同步 API请求
// alibaba.campus.guard.controller.configsync
//
// 门禁控制器配置项同步
type AlibabaCampusGuardControllerConfigsyncAPIRequest struct {
	model.Params
	// 查询参数类
	_controllerQuery *ControllerQuery
}

// NewAlibabaCampusGuardControllerConfigsyncRequest 初始化AlibabaCampusGuardControllerConfigsyncAPIRequest对象
func NewAlibabaCampusGuardControllerConfigsyncRequest() *AlibabaCampusGuardControllerConfigsyncAPIRequest {
	return &AlibabaCampusGuardControllerConfigsyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusGuardControllerConfigsyncAPIRequest) Reset() {
	r._controllerQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusGuardControllerConfigsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.guard.controller.configsync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusGuardControllerConfigsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusGuardControllerConfigsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetControllerQuery is ControllerQuery Setter
// 查询参数类
func (r *AlibabaCampusGuardControllerConfigsyncAPIRequest) SetControllerQuery(_controllerQuery *ControllerQuery) error {
	r._controllerQuery = _controllerQuery
	r.Set("controller_query", _controllerQuery)
	return nil
}

// GetControllerQuery ControllerQuery Getter
func (r AlibabaCampusGuardControllerConfigsyncAPIRequest) GetControllerQuery() *ControllerQuery {
	return r._controllerQuery
}

var poolAlibabaCampusGuardControllerConfigsyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusGuardControllerConfigsyncRequest()
	},
}

// GetAlibabaCampusGuardControllerConfigsyncRequest 从 sync.Pool 获取 AlibabaCampusGuardControllerConfigsyncAPIRequest
func GetAlibabaCampusGuardControllerConfigsyncAPIRequest() *AlibabaCampusGuardControllerConfigsyncAPIRequest {
	return poolAlibabaCampusGuardControllerConfigsyncAPIRequest.Get().(*AlibabaCampusGuardControllerConfigsyncAPIRequest)
}

// ReleaseAlibabaCampusGuardControllerConfigsyncAPIRequest 将 AlibabaCampusGuardControllerConfigsyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusGuardControllerConfigsyncAPIRequest(v *AlibabaCampusGuardControllerConfigsyncAPIRequest) {
	v.Reset()
	poolAlibabaCampusGuardControllerConfigsyncAPIRequest.Put(v)
}
