package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusGuardControllerOfflinelogAPIRequest 门禁控制器离线日志同步 API请求
// alibaba.campus.guard.controller.offlinelog
//
// 门禁控制器离线日志同步
type AlibabaCampusGuardControllerOfflinelogAPIRequest struct {
	model.Params
	// 离线数据DTO
	_controllerOfflineRequestDto *ControllerOfflineRequestDto
}

// NewAlibabaCampusGuardControllerOfflinelogRequest 初始化AlibabaCampusGuardControllerOfflinelogAPIRequest对象
func NewAlibabaCampusGuardControllerOfflinelogRequest() *AlibabaCampusGuardControllerOfflinelogAPIRequest {
	return &AlibabaCampusGuardControllerOfflinelogAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusGuardControllerOfflinelogAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.guard.controller.offlinelog"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusGuardControllerOfflinelogAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusGuardControllerOfflinelogAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetControllerOfflineRequestDto is ControllerOfflineRequestDto Setter
// 离线数据DTO
func (r *AlibabaCampusGuardControllerOfflinelogAPIRequest) SetControllerOfflineRequestDto(_controllerOfflineRequestDto *ControllerOfflineRequestDto) error {
	r._controllerOfflineRequestDto = _controllerOfflineRequestDto
	r.Set("controller_offline_request_dto", _controllerOfflineRequestDto)
	return nil
}

// GetControllerOfflineRequestDto ControllerOfflineRequestDto Getter
func (r AlibabaCampusGuardControllerOfflinelogAPIRequest) GetControllerOfflineRequestDto() *ControllerOfflineRequestDto {
	return r._controllerOfflineRequestDto
}
