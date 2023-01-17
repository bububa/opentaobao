package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayCreativeVideoUnbindAPIRequest 创意与视频解绑接口 API请求
// taobao.subway.creative.video.unbind
//
// 将创意与视频解绑
type TaobaoSubwayCreativeVideoUnbindAPIRequest struct {
	model.Params
	// 淘宝账号昵称
	_nick string
	// 视频ID
	_videoId int64
	// 创意ID
	_creativeId int64
}

// NewTaobaoSubwayCreativeVideoUnbindRequest 初始化TaobaoSubwayCreativeVideoUnbindAPIRequest对象
func NewTaobaoSubwayCreativeVideoUnbindRequest() *TaobaoSubwayCreativeVideoUnbindAPIRequest {
	return &TaobaoSubwayCreativeVideoUnbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubwayCreativeVideoUnbindAPIRequest) GetApiMethodName() string {
	return "taobao.subway.creative.video.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubwayCreativeVideoUnbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSubwayCreativeVideoUnbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 淘宝账号昵称
func (r *TaobaoSubwayCreativeVideoUnbindAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSubwayCreativeVideoUnbindAPIRequest) GetNick() string {
	return r._nick
}

// SetVideoId is VideoId Setter
// 视频ID
func (r *TaobaoSubwayCreativeVideoUnbindAPIRequest) SetVideoId(_videoId int64) error {
	r._videoId = _videoId
	r.Set("video_id", _videoId)
	return nil
}

// GetVideoId VideoId Getter
func (r TaobaoSubwayCreativeVideoUnbindAPIRequest) GetVideoId() int64 {
	return r._videoId
}

// SetCreativeId is CreativeId Setter
// 创意ID
func (r *TaobaoSubwayCreativeVideoUnbindAPIRequest) SetCreativeId(_creativeId int64) error {
	r._creativeId = _creativeId
	r.Set("creative_id", _creativeId)
	return nil
}

// GetCreativeId CreativeId Getter
func (r TaobaoSubwayCreativeVideoUnbindAPIRequest) GetCreativeId() int64 {
	return r._creativeId
}
