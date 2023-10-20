package mirage

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuMirageQueryPermissionAPIRequest 优酷播控查询是否可播API API请求
// youku.mirage.query.permission
//
// 根据节目ID或者VID查询视频或者节目是否可以播放
type YoukuMirageQueryPermissionAPIRequest struct {
	model.Params
	// 入参
	_permissionRequestDto *PermissionRequestDto
}

// NewYoukuMirageQueryPermissionRequest 初始化YoukuMirageQueryPermissionAPIRequest对象
func NewYoukuMirageQueryPermissionRequest() *YoukuMirageQueryPermissionAPIRequest {
	return &YoukuMirageQueryPermissionAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YoukuMirageQueryPermissionAPIRequest) Reset() {
	r._permissionRequestDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuMirageQueryPermissionAPIRequest) GetApiMethodName() string {
	return "youku.mirage.query.permission"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuMirageQueryPermissionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuMirageQueryPermissionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPermissionRequestDto is PermissionRequestDto Setter
// 入参
func (r *YoukuMirageQueryPermissionAPIRequest) SetPermissionRequestDto(_permissionRequestDto *PermissionRequestDto) error {
	r._permissionRequestDto = _permissionRequestDto
	r.Set("permission_request_dto", _permissionRequestDto)
	return nil
}

// GetPermissionRequestDto PermissionRequestDto Getter
func (r YoukuMirageQueryPermissionAPIRequest) GetPermissionRequestDto() *PermissionRequestDto {
	return r._permissionRequestDto
}

var poolYoukuMirageQueryPermissionAPIRequest = sync.Pool{
	New: func() any {
		return NewYoukuMirageQueryPermissionRequest()
	},
}

// GetYoukuMirageQueryPermissionRequest 从 sync.Pool 获取 YoukuMirageQueryPermissionAPIRequest
func GetYoukuMirageQueryPermissionAPIRequest() *YoukuMirageQueryPermissionAPIRequest {
	return poolYoukuMirageQueryPermissionAPIRequest.Get().(*YoukuMirageQueryPermissionAPIRequest)
}

// ReleaseYoukuMirageQueryPermissionAPIRequest 将 YoukuMirageQueryPermissionAPIRequest 放入 sync.Pool
func ReleaseYoukuMirageQueryPermissionAPIRequest(v *YoukuMirageQueryPermissionAPIRequest) {
	v.Reset()
	poolYoukuMirageQueryPermissionAPIRequest.Put(v)
}
