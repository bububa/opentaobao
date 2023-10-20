package taotv

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaotvCarouselPlaylistGetAPIRequest 根据频道ID获取频道下节目单以及当前播放 API请求
// taobao.taotv.carousel.playlist.get
//
// 根据频道ID获取频道下节目单以及当前播放，包括所有视频源的视频
type TaobaoTaotvCarouselPlaylistGetAPIRequest struct {
	model.Params
	// 设备信息
	_systemInfo string
	// 频道ID
	_channelId int64
}

// NewTaobaoTaotvCarouselPlaylistGetRequest 初始化TaobaoTaotvCarouselPlaylistGetAPIRequest对象
func NewTaobaoTaotvCarouselPlaylistGetRequest() *TaobaoTaotvCarouselPlaylistGetAPIRequest {
	return &TaobaoTaotvCarouselPlaylistGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTaotvCarouselPlaylistGetAPIRequest) Reset() {
	r._systemInfo = ""
	r._channelId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTaotvCarouselPlaylistGetAPIRequest) GetApiMethodName() string {
	return "taobao.taotv.carousel.playlist.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTaotvCarouselPlaylistGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTaotvCarouselPlaylistGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSystemInfo is SystemInfo Setter
// 设备信息
func (r *TaobaoTaotvCarouselPlaylistGetAPIRequest) SetSystemInfo(_systemInfo string) error {
	r._systemInfo = _systemInfo
	r.Set("system_info", _systemInfo)
	return nil
}

// GetSystemInfo SystemInfo Getter
func (r TaobaoTaotvCarouselPlaylistGetAPIRequest) GetSystemInfo() string {
	return r._systemInfo
}

// SetChannelId is ChannelId Setter
// 频道ID
func (r *TaobaoTaotvCarouselPlaylistGetAPIRequest) SetChannelId(_channelId int64) error {
	r._channelId = _channelId
	r.Set("channel_id", _channelId)
	return nil
}

// GetChannelId ChannelId Getter
func (r TaobaoTaotvCarouselPlaylistGetAPIRequest) GetChannelId() int64 {
	return r._channelId
}

var poolTaobaoTaotvCarouselPlaylistGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTaotvCarouselPlaylistGetRequest()
	},
}

// GetTaobaoTaotvCarouselPlaylistGetRequest 从 sync.Pool 获取 TaobaoTaotvCarouselPlaylistGetAPIRequest
func GetTaobaoTaotvCarouselPlaylistGetAPIRequest() *TaobaoTaotvCarouselPlaylistGetAPIRequest {
	return poolTaobaoTaotvCarouselPlaylistGetAPIRequest.Get().(*TaobaoTaotvCarouselPlaylistGetAPIRequest)
}

// ReleaseTaobaoTaotvCarouselPlaylistGetAPIRequest 将 TaobaoTaotvCarouselPlaylistGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTaotvCarouselPlaylistGetAPIRequest(v *TaobaoTaotvCarouselPlaylistGetAPIRequest) {
	v.Reset()
	poolTaobaoTaotvCarouselPlaylistGetAPIRequest.Put(v)
}
