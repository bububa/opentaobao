package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomevideosyncAPIRequest 视频草稿信息同步 API请求
// alibaba.alihouse.newhome.video.sync
//
// 接收视频信息记录
type AlibabaalihousenewhomevideosyncAPIRequest struct {
	model.Params
	// 草稿对下
	_video *VideoDraftDto
}

// NewAlibabaalihousenewhomevideosyncRequest 初始化AlibabaalihousenewhomevideosyncAPIRequest对象
func NewAlibabaalihousenewhomevideosyncRequest() *AlibabaalihousenewhomevideosyncAPIRequest {
	return &AlibabaalihousenewhomevideosyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomevideosyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.video.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomevideosyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomevideosyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVideo is Video Setter
// 草稿对下
func (r *AlibabaalihousenewhomevideosyncAPIRequest) SetVideo(_video *VideoDraftDto) error {
	r._video = _video
	r.Set("video", _video)
	return nil
}

// GetVideo Video Getter
func (r AlibabaalihousenewhomevideosyncAPIRequest) GetVideo() *VideoDraftDto {
	return r._video
}
