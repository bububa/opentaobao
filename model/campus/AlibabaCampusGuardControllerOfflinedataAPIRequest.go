package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusGuardControllerOfflinedataAPIRequest 点位离线数据拉取 API请求
// alibaba.campus.guard.controller.offlinedata
//
// 点位离线数据拉取
type AlibabaCampusGuardControllerOfflinedataAPIRequest struct {
	model.Params
	// requestParam
	_controllerOfflineRequestDto *ControllerOfflineRequestDto
}

// NewAlibabaCampusGuardControllerOfflinedataRequest 初始化AlibabaCampusGuardControllerOfflinedataAPIRequest对象
func NewAlibabaCampusGuardControllerOfflinedataRequest() *AlibabaCampusGuardControllerOfflinedataAPIRequest {
	return &AlibabaCampusGuardControllerOfflinedataAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusGuardControllerOfflinedataAPIRequest) Reset() {
	r._controllerOfflineRequestDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusGuardControllerOfflinedataAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.guard.controller.offlinedata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusGuardControllerOfflinedataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusGuardControllerOfflinedataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetControllerOfflineRequestDto is ControllerOfflineRequestDto Setter
// requestParam
func (r *AlibabaCampusGuardControllerOfflinedataAPIRequest) SetControllerOfflineRequestDto(_controllerOfflineRequestDto *ControllerOfflineRequestDto) error {
	r._controllerOfflineRequestDto = _controllerOfflineRequestDto
	r.Set("controller_offline_request_dto", _controllerOfflineRequestDto)
	return nil
}

// GetControllerOfflineRequestDto ControllerOfflineRequestDto Getter
func (r AlibabaCampusGuardControllerOfflinedataAPIRequest) GetControllerOfflineRequestDto() *ControllerOfflineRequestDto {
	return r._controllerOfflineRequestDto
}

var poolAlibabaCampusGuardControllerOfflinedataAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusGuardControllerOfflinedataRequest()
	},
}

// GetAlibabaCampusGuardControllerOfflinedataRequest 从 sync.Pool 获取 AlibabaCampusGuardControllerOfflinedataAPIRequest
func GetAlibabaCampusGuardControllerOfflinedataAPIRequest() *AlibabaCampusGuardControllerOfflinedataAPIRequest {
	return poolAlibabaCampusGuardControllerOfflinedataAPIRequest.Get().(*AlibabaCampusGuardControllerOfflinedataAPIRequest)
}

// ReleaseAlibabaCampusGuardControllerOfflinedataAPIRequest 将 AlibabaCampusGuardControllerOfflinedataAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusGuardControllerOfflinedataAPIRequest(v *AlibabaCampusGuardControllerOfflinedataAPIRequest) {
	v.Reset()
	poolAlibabaCampusGuardControllerOfflinedataAPIRequest.Put(v)
}
