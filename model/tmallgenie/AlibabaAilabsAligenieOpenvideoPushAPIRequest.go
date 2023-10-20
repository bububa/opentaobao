package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsaligenieopenvideopushAPIRequest 天猫精灵内容库视频分集数据推送接口 API请求
// alibaba.ailabs.aligenie.openvideo.push
//
// 天猫精灵内容库视频分集数据推送接口
type AlibabaailabsaligenieopenvideopushAPIRequest struct {
	model.Params
	// 待推送的视频数据
	_videos []RawSingleVideo
}

// NewAlibabaailabsaligenieopenvideopushRequest 初始化AlibabaailabsaligenieopenvideopushAPIRequest对象
func NewAlibabaailabsaligenieopenvideopushRequest() *AlibabaailabsaligenieopenvideopushAPIRequest {
	return &AlibabaailabsaligenieopenvideopushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabsaligenieopenvideopushAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.aligenie.openvideo.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabsaligenieopenvideopushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabsaligenieopenvideopushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVideos is Videos Setter
// 待推送的视频数据
func (r *AlibabaailabsaligenieopenvideopushAPIRequest) SetVideos(_videos []RawSingleVideo) error {
	r._videos = _videos
	r.Set("videos", _videos)
	return nil
}

// GetVideos Videos Getter
func (r AlibabaailabsaligenieopenvideopushAPIRequest) GetVideos() []RawSingleVideo {
	return r._videos
}
