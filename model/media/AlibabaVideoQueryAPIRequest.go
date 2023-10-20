package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabavideoqueryAPIRequest 查询视频信息 API请求
// alibaba.video.query
//
// 查询视频信息
type AlibabavideoqueryAPIRequest struct {
	model.Params
	// 视频id
	_videoId int64
}

// NewAlibabavideoqueryRequest 初始化AlibabavideoqueryAPIRequest对象
func NewAlibabavideoqueryRequest() *AlibabavideoqueryAPIRequest {
	return &AlibabavideoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabavideoqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.video.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabavideoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabavideoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVideoId is VideoId Setter
// 视频id
func (r *AlibabavideoqueryAPIRequest) SetVideoId(_videoId int64) error {
	r._videoId = _videoId
	r.Set("video_id", _videoId)
	return nil
}

// GetVideoId VideoId Getter
func (r AlibabavideoqueryAPIRequest) GetVideoId() int64 {
	return r._videoId
}
