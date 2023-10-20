package media

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaVideoQueryAPIRequest 查询视频信息 API请求
// alibaba.video.query
//
// 查询视频信息
type AlibabaVideoQueryAPIRequest struct {
	model.Params
	// 视频id
	_videoId int64
}

// NewAlibabaVideoQueryRequest 初始化AlibabaVideoQueryAPIRequest对象
func NewAlibabaVideoQueryRequest() *AlibabaVideoQueryAPIRequest {
	return &AlibabaVideoQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaVideoQueryAPIRequest) Reset() {
	r._videoId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaVideoQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.video.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaVideoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaVideoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVideoId is VideoId Setter
// 视频id
func (r *AlibabaVideoQueryAPIRequest) SetVideoId(_videoId int64) error {
	r._videoId = _videoId
	r.Set("video_id", _videoId)
	return nil
}

// GetVideoId VideoId Getter
func (r AlibabaVideoQueryAPIRequest) GetVideoId() int64 {
	return r._videoId
}

var poolAlibabaVideoQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaVideoQueryRequest()
	},
}

// GetAlibabaVideoQueryRequest 从 sync.Pool 获取 AlibabaVideoQueryAPIRequest
func GetAlibabaVideoQueryAPIRequest() *AlibabaVideoQueryAPIRequest {
	return poolAlibabaVideoQueryAPIRequest.Get().(*AlibabaVideoQueryAPIRequest)
}

// ReleaseAlibabaVideoQueryAPIRequest 将 AlibabaVideoQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaVideoQueryAPIRequest(v *AlibabaVideoQueryAPIRequest) {
	v.Reset()
	poolAlibabaVideoQueryAPIRequest.Put(v)
}
