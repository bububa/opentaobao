package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubwayvideostategetAPIRequest 获取视频状态 API请求
// taobao.subway.video.state.get
//
// 获取已上传视频的状态
type TaobaosubwayvideostategetAPIRequest struct {
	model.Params
	// 账号昵称
	_nick string
	// videoId
	_videoId int64
}

// NewTaobaosubwayvideostategetRequest 初始化TaobaosubwayvideostategetAPIRequest对象
func NewTaobaosubwayvideostategetRequest() *TaobaosubwayvideostategetAPIRequest {
	return &TaobaosubwayvideostategetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosubwayvideostategetAPIRequest) GetApiMethodName() string {
	return "taobao.subway.video.state.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosubwayvideostategetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosubwayvideostategetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 账号昵称
func (r *TaobaosubwayvideostategetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosubwayvideostategetAPIRequest) GetNick() string {
	return r._nick
}

// SetVideoId is VideoId Setter
// videoId
func (r *TaobaosubwayvideostategetAPIRequest) SetVideoId(_videoId int64) error {
	r._videoId = _videoId
	r.Set("video_id", _videoId)
	return nil
}

// GetVideoId VideoId Getter
func (r TaobaosubwayvideostategetAPIRequest) GetVideoId() int64 {
	return r._videoId
}
