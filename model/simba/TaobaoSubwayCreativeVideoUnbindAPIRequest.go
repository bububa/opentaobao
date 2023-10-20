package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubwaycreativevideounbindAPIRequest 创意与视频解绑接口 API请求
// taobao.subway.creative.video.unbind
//
// 将创意与视频解绑
type TaobaosubwaycreativevideounbindAPIRequest struct {
	model.Params
	// 淘宝账号昵称
	_nick string
	// 视频ID
	_videoId int64
	// 创意ID
	_creativeId int64
}

// NewTaobaosubwaycreativevideounbindRequest 初始化TaobaosubwaycreativevideounbindAPIRequest对象
func NewTaobaosubwaycreativevideounbindRequest() *TaobaosubwaycreativevideounbindAPIRequest {
	return &TaobaosubwaycreativevideounbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosubwaycreativevideounbindAPIRequest) GetApiMethodName() string {
	return "taobao.subway.creative.video.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosubwaycreativevideounbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosubwaycreativevideounbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 淘宝账号昵称
func (r *TaobaosubwaycreativevideounbindAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosubwaycreativevideounbindAPIRequest) GetNick() string {
	return r._nick
}

// SetVideoId is VideoId Setter
// 视频ID
func (r *TaobaosubwaycreativevideounbindAPIRequest) SetVideoId(_videoId int64) error {
	r._videoId = _videoId
	r.Set("video_id", _videoId)
	return nil
}

// GetVideoId VideoId Getter
func (r TaobaosubwaycreativevideounbindAPIRequest) GetVideoId() int64 {
	return r._videoId
}

// SetCreativeId is CreativeId Setter
// 创意ID
func (r *TaobaosubwaycreativevideounbindAPIRequest) SetCreativeId(_creativeId int64) error {
	r._creativeId = _creativeId
	r.Set("creative_id", _creativeId)
	return nil
}

// GetCreativeId CreativeId Getter
func (r TaobaosubwaycreativevideounbindAPIRequest) GetCreativeId() int64 {
	return r._creativeId
}
