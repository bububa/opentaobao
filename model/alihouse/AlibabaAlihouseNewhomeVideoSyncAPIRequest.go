package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeVideoSyncAPIRequest 视频草稿信息同步 API请求
// alibaba.alihouse.newhome.video.sync
//
// 接收视频信息记录
type AlibabaAlihouseNewhomeVideoSyncAPIRequest struct {
	model.Params
	// 草稿对下
	_video *VideoDraftDto
}

// NewAlibabaAlihouseNewhomeVideoSyncRequest 初始化AlibabaAlihouseNewhomeVideoSyncAPIRequest对象
func NewAlibabaAlihouseNewhomeVideoSyncRequest() *AlibabaAlihouseNewhomeVideoSyncAPIRequest {
	return &AlibabaAlihouseNewhomeVideoSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeVideoSyncAPIRequest) Reset() {
	r._video = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeVideoSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.video.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeVideoSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeVideoSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVideo is Video Setter
// 草稿对下
func (r *AlibabaAlihouseNewhomeVideoSyncAPIRequest) SetVideo(_video *VideoDraftDto) error {
	r._video = _video
	r.Set("video", _video)
	return nil
}

// GetVideo Video Getter
func (r AlibabaAlihouseNewhomeVideoSyncAPIRequest) GetVideo() *VideoDraftDto {
	return r._video
}

var poolAlibabaAlihouseNewhomeVideoSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeVideoSyncRequest()
	},
}

// GetAlibabaAlihouseNewhomeVideoSyncRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeVideoSyncAPIRequest
func GetAlibabaAlihouseNewhomeVideoSyncAPIRequest() *AlibabaAlihouseNewhomeVideoSyncAPIRequest {
	return poolAlibabaAlihouseNewhomeVideoSyncAPIRequest.Get().(*AlibabaAlihouseNewhomeVideoSyncAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeVideoSyncAPIRequest 将 AlibabaAlihouseNewhomeVideoSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeVideoSyncAPIRequest(v *AlibabaAlihouseNewhomeVideoSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeVideoSyncAPIRequest.Put(v)
}
