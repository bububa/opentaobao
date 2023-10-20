package taotv

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaotvVideoPlaylistPageAPIRequest 分页获取所有播单 API请求
// taobao.taotv.video.playlist.page
//
// 获取所有播单信息（分页）
type TaobaoTaotvVideoPlaylistPageAPIRequest struct {
	model.Params
	// 客户端信息
	_systemInfo string
	// 当前页（从1开始）
	_pageNo int64
}

// NewTaobaoTaotvVideoPlaylistPageRequest 初始化TaobaoTaotvVideoPlaylistPageAPIRequest对象
func NewTaobaoTaotvVideoPlaylistPageRequest() *TaobaoTaotvVideoPlaylistPageAPIRequest {
	return &TaobaoTaotvVideoPlaylistPageAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTaotvVideoPlaylistPageAPIRequest) Reset() {
	r._systemInfo = ""
	r._pageNo = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTaotvVideoPlaylistPageAPIRequest) GetApiMethodName() string {
	return "taobao.taotv.video.playlist.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTaotvVideoPlaylistPageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTaotvVideoPlaylistPageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSystemInfo is SystemInfo Setter
// 客户端信息
func (r *TaobaoTaotvVideoPlaylistPageAPIRequest) SetSystemInfo(_systemInfo string) error {
	r._systemInfo = _systemInfo
	r.Set("system_info", _systemInfo)
	return nil
}

// GetSystemInfo SystemInfo Getter
func (r TaobaoTaotvVideoPlaylistPageAPIRequest) GetSystemInfo() string {
	return r._systemInfo
}

// SetPageNo is PageNo Setter
// 当前页（从1开始）
func (r *TaobaoTaotvVideoPlaylistPageAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoTaotvVideoPlaylistPageAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

var poolTaobaoTaotvVideoPlaylistPageAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTaotvVideoPlaylistPageRequest()
	},
}

// GetTaobaoTaotvVideoPlaylistPageRequest 从 sync.Pool 获取 TaobaoTaotvVideoPlaylistPageAPIRequest
func GetTaobaoTaotvVideoPlaylistPageAPIRequest() *TaobaoTaotvVideoPlaylistPageAPIRequest {
	return poolTaobaoTaotvVideoPlaylistPageAPIRequest.Get().(*TaobaoTaotvVideoPlaylistPageAPIRequest)
}

// ReleaseTaobaoTaotvVideoPlaylistPageAPIRequest 将 TaobaoTaotvVideoPlaylistPageAPIRequest 放入 sync.Pool
func ReleaseTaobaoTaotvVideoPlaylistPageAPIRequest(v *TaobaoTaotvVideoPlaylistPageAPIRequest) {
	v.Reset()
	poolTaobaoTaotvVideoPlaylistPageAPIRequest.Put(v)
}
