package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusguardcontrollerconfigsyncAPIRequest 门禁控制器配置项同步 API请求
// alibaba.campus.guard.controller.configsync
//
// 门禁控制器配置项同步
type AlibabacampusguardcontrollerconfigsyncAPIRequest struct {
	model.Params
	// 查询参数类
	_controllerQuery *ControllerQuery
}

// NewAlibabacampusguardcontrollerconfigsyncRequest 初始化AlibabacampusguardcontrollerconfigsyncAPIRequest对象
func NewAlibabacampusguardcontrollerconfigsyncRequest() *AlibabacampusguardcontrollerconfigsyncAPIRequest {
	return &AlibabacampusguardcontrollerconfigsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusguardcontrollerconfigsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.guard.controller.configsync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusguardcontrollerconfigsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusguardcontrollerconfigsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetControllerQuery is ControllerQuery Setter
// 查询参数类
func (r *AlibabacampusguardcontrollerconfigsyncAPIRequest) SetControllerQuery(_controllerQuery *ControllerQuery) error {
	r._controllerQuery = _controllerQuery
	r.Set("controller_query", _controllerQuery)
	return nil
}

// GetControllerQuery ControllerQuery Getter
func (r AlibabacampusguardcontrollerconfigsyncAPIRequest) GetControllerQuery() *ControllerQuery {
	return r._controllerQuery
}
