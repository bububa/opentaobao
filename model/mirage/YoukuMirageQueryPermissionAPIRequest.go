package mirage

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukumiragequerypermissionAPIRequest 优酷播控查询是否可播API API请求
// youku.mirage.query.permission
//
// 根据节目ID或者VID查询视频或者节目是否可以播放
type YoukumiragequerypermissionAPIRequest struct {
	model.Params
	// 入参
	_permissionRequestDto *PermissionRequestDto
}

// NewYoukumiragequerypermissionRequest 初始化YoukumiragequerypermissionAPIRequest对象
func NewYoukumiragequerypermissionRequest() *YoukumiragequerypermissionAPIRequest {
	return &YoukumiragequerypermissionAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukumiragequerypermissionAPIRequest) GetApiMethodName() string {
	return "youku.mirage.query.permission"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukumiragequerypermissionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukumiragequerypermissionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPermissionRequestDto is PermissionRequestDto Setter
// 入参
func (r *YoukumiragequerypermissionAPIRequest) SetPermissionRequestDto(_permissionRequestDto *PermissionRequestDto) error {
	r._permissionRequestDto = _permissionRequestDto
	r.Set("permission_request_dto", _permissionRequestDto)
	return nil
}

// GetPermissionRequestDto PermissionRequestDto Getter
func (r YoukumiragequerypermissionAPIRequest) GetPermissionRequestDto() *PermissionRequestDto {
	return r._permissionRequestDto
}
