package campus

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusGuardControllerOfflinelogAPIRequest) Reset() {
	r._controllerOfflineRequestDto = nil
	r.Params.ToZero()
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

var poolAlibabaCampusGuardControllerOfflinelogAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusGuardControllerOfflinelogRequest()
	},
}

// GetAlibabaCampusGuardControllerOfflinelogRequest 从 sync.Pool 获取 AlibabaCampusGuardControllerOfflinelogAPIRequest
func GetAlibabaCampusGuardControllerOfflinelogAPIRequest() *AlibabaCampusGuardControllerOfflinelogAPIRequest {
	return poolAlibabaCampusGuardControllerOfflinelogAPIRequest.Get().(*AlibabaCampusGuardControllerOfflinelogAPIRequest)
}

// ReleaseAlibabaCampusGuardControllerOfflinelogAPIRequest 将 AlibabaCampusGuardControllerOfflinelogAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusGuardControllerOfflinelogAPIRequest(v *AlibabaCampusGuardControllerOfflinelogAPIRequest) {
	v.Reset()
	poolAlibabaCampusGuardControllerOfflinelogAPIRequest.Put(v)
}
