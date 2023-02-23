package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieOpenvideoPushAPIRequest 天猫精灵内容库视频分集数据推送接口 API请求
// alibaba.ailabs.aligenie.openvideo.push
//
// 天猫精灵内容库视频分集数据推送接口
type AlibabaAilabsAligenieOpenvideoPushAPIRequest struct {
	model.Params
	// 待推送的视频数据
	_videos []RawSingleVideo
}

// NewAlibabaAilabsAligenieOpenvideoPushRequest 初始化AlibabaAilabsAligenieOpenvideoPushAPIRequest对象
func NewAlibabaAilabsAligenieOpenvideoPushRequest() *AlibabaAilabsAligenieOpenvideoPushAPIRequest {
	return &AlibabaAilabsAligenieOpenvideoPushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieOpenvideoPushAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.aligenie.openvideo.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieOpenvideoPushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsAligenieOpenvideoPushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVideos is Videos Setter
// 待推送的视频数据
func (r *AlibabaAilabsAligenieOpenvideoPushAPIRequest) SetVideos(_videos []RawSingleVideo) error {
	r._videos = _videos
	r.Set("videos", _videos)
	return nil
}

// GetVideos Videos Getter
func (r AlibabaAilabsAligenieOpenvideoPushAPIRequest) GetVideos() []RawSingleVideo {
	return r._videos
}
