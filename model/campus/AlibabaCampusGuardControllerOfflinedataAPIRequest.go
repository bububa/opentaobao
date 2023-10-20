package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusguardcontrollerofflinedataAPIRequest 点位离线数据拉取 API请求
// alibaba.campus.guard.controller.offlinedata
//
// 点位离线数据拉取
type AlibabacampusguardcontrollerofflinedataAPIRequest struct {
	model.Params
	// requestParam
	_controllerOfflineRequestDto *ControllerOfflineRequestDto
}

// NewAlibabacampusguardcontrollerofflinedataRequest 初始化AlibabacampusguardcontrollerofflinedataAPIRequest对象
func NewAlibabacampusguardcontrollerofflinedataRequest() *AlibabacampusguardcontrollerofflinedataAPIRequest {
	return &AlibabacampusguardcontrollerofflinedataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusguardcontrollerofflinedataAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.guard.controller.offlinedata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusguardcontrollerofflinedataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusguardcontrollerofflinedataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetControllerOfflineRequestDto is ControllerOfflineRequestDto Setter
// requestParam
func (r *AlibabacampusguardcontrollerofflinedataAPIRequest) SetControllerOfflineRequestDto(_controllerOfflineRequestDto *ControllerOfflineRequestDto) error {
	r._controllerOfflineRequestDto = _controllerOfflineRequestDto
	r.Set("controller_offline_request_dto", _controllerOfflineRequestDto)
	return nil
}

// GetControllerOfflineRequestDto ControllerOfflineRequestDto Getter
func (r AlibabacampusguardcontrollerofflinedataAPIRequest) GetControllerOfflineRequestDto() *ControllerOfflineRequestDto {
	return r._controllerOfflineRequestDto
}
