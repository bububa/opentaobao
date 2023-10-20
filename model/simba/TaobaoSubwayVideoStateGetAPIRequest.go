package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayVideoStateGetAPIRequest 获取视频状态 API请求
// taobao.subway.video.state.get
//
// 获取已上传视频的状态
type TaobaoSubwayVideoStateGetAPIRequest struct {
	model.Params
	// 账号昵称
	_nick string
	// videoId
	_videoId int64
}

// NewTaobaoSubwayVideoStateGetRequest 初始化TaobaoSubwayVideoStateGetAPIRequest对象
func NewTaobaoSubwayVideoStateGetRequest() *TaobaoSubwayVideoStateGetAPIRequest {
	return &TaobaoSubwayVideoStateGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSubwayVideoStateGetAPIRequest) Reset() {
	r._nick = ""
	r._videoId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubwayVideoStateGetAPIRequest) GetApiMethodName() string {
	return "taobao.subway.video.state.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubwayVideoStateGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSubwayVideoStateGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 账号昵称
func (r *TaobaoSubwayVideoStateGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSubwayVideoStateGetAPIRequest) GetNick() string {
	return r._nick
}

// SetVideoId is VideoId Setter
// videoId
func (r *TaobaoSubwayVideoStateGetAPIRequest) SetVideoId(_videoId int64) error {
	r._videoId = _videoId
	r.Set("video_id", _videoId)
	return nil
}

// GetVideoId VideoId Getter
func (r TaobaoSubwayVideoStateGetAPIRequest) GetVideoId() int64 {
	return r._videoId
}

var poolTaobaoSubwayVideoStateGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSubwayVideoStateGetRequest()
	},
}

// GetTaobaoSubwayVideoStateGetRequest 从 sync.Pool 获取 TaobaoSubwayVideoStateGetAPIRequest
func GetTaobaoSubwayVideoStateGetAPIRequest() *TaobaoSubwayVideoStateGetAPIRequest {
	return poolTaobaoSubwayVideoStateGetAPIRequest.Get().(*TaobaoSubwayVideoStateGetAPIRequest)
}

// ReleaseTaobaoSubwayVideoStateGetAPIRequest 将 TaobaoSubwayVideoStateGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSubwayVideoStateGetAPIRequest(v *TaobaoSubwayVideoStateGetAPIRequest) {
	v.Reset()
	poolTaobaoSubwayVideoStateGetAPIRequest.Put(v)
}
