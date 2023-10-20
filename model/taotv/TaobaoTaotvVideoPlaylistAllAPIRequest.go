package taotv

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaotvVideoPlaylistAllAPIRequest 获取播单列表 API请求
// taobao.taotv.video.playlist.all
//
// 根据牌照和视频源等获取播单列表
type TaobaoTaotvVideoPlaylistAllAPIRequest struct {
	model.Params
	// 系统信息
	_systemInfo string
}

// NewTaobaoTaotvVideoPlaylistAllRequest 初始化TaobaoTaotvVideoPlaylistAllAPIRequest对象
func NewTaobaoTaotvVideoPlaylistAllRequest() *TaobaoTaotvVideoPlaylistAllAPIRequest {
	return &TaobaoTaotvVideoPlaylistAllAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTaotvVideoPlaylistAllAPIRequest) Reset() {
	r._systemInfo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTaotvVideoPlaylistAllAPIRequest) GetApiMethodName() string {
	return "taobao.taotv.video.playlist.all"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTaotvVideoPlaylistAllAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTaotvVideoPlaylistAllAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSystemInfo is SystemInfo Setter
// 系统信息
func (r *TaobaoTaotvVideoPlaylistAllAPIRequest) SetSystemInfo(_systemInfo string) error {
	r._systemInfo = _systemInfo
	r.Set("system_info", _systemInfo)
	return nil
}

// GetSystemInfo SystemInfo Getter
func (r TaobaoTaotvVideoPlaylistAllAPIRequest) GetSystemInfo() string {
	return r._systemInfo
}

var poolTaobaoTaotvVideoPlaylistAllAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTaotvVideoPlaylistAllRequest()
	},
}

// GetTaobaoTaotvVideoPlaylistAllRequest 从 sync.Pool 获取 TaobaoTaotvVideoPlaylistAllAPIRequest
func GetTaobaoTaotvVideoPlaylistAllAPIRequest() *TaobaoTaotvVideoPlaylistAllAPIRequest {
	return poolTaobaoTaotvVideoPlaylistAllAPIRequest.Get().(*TaobaoTaotvVideoPlaylistAllAPIRequest)
}

// ReleaseTaobaoTaotvVideoPlaylistAllAPIRequest 将 TaobaoTaotvVideoPlaylistAllAPIRequest 放入 sync.Pool
func ReleaseTaobaoTaotvVideoPlaylistAllAPIRequest(v *TaobaoTaotvVideoPlaylistAllAPIRequest) {
	v.Reset()
	poolTaobaoTaotvVideoPlaylistAllAPIRequest.Put(v)
}
