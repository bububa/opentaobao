package wenyuvideo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukuwenyuvideoseetagetAPIRequest 只看TA API请求
// youku.wenyuvideo.seeta.get
//
// 只看Ta对外输出
type YoukuwenyuvideoseetagetAPIRequest struct {
	model.Params
	// 视频字符串形式id
	_videoStrId string
}

// NewYoukuwenyuvideoseetagetRequest 初始化YoukuwenyuvideoseetagetAPIRequest对象
func NewYoukuwenyuvideoseetagetRequest() *YoukuwenyuvideoseetagetAPIRequest {
	return &YoukuwenyuvideoseetagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuwenyuvideoseetagetAPIRequest) GetApiMethodName() string {
	return "youku.wenyuvideo.seeta.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuwenyuvideoseetagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuwenyuvideoseetagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVideoStrId is VideoStrId Setter
// 视频字符串形式id
func (r *YoukuwenyuvideoseetagetAPIRequest) SetVideoStrId(_videoStrId string) error {
	r._videoStrId = _videoStrId
	r.Set("video_str_id", _videoStrId)
	return nil
}

// GetVideoStrId VideoStrId Getter
func (r YoukuwenyuvideoseetagetAPIRequest) GetVideoStrId() string {
	return r._videoStrId
}
