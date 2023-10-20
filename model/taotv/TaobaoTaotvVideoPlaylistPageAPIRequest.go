package taotv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotaotvvideoplaylistpageAPIRequest 分页获取所有播单 API请求
// taobao.taotv.video.playlist.page
//
// 获取所有播单信息（分页）
type TaobaotaotvvideoplaylistpageAPIRequest struct {
	model.Params
	// 客户端信息
	_systemInfo string
	// 当前页（从1开始）
	_pageNo int64
}

// NewTaobaotaotvvideoplaylistpageRequest 初始化TaobaotaotvvideoplaylistpageAPIRequest对象
func NewTaobaotaotvvideoplaylistpageRequest() *TaobaotaotvvideoplaylistpageAPIRequest {
	return &TaobaotaotvvideoplaylistpageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotaotvvideoplaylistpageAPIRequest) GetApiMethodName() string {
	return "taobao.taotv.video.playlist.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotaotvvideoplaylistpageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotaotvvideoplaylistpageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSystemInfo is SystemInfo Setter
// 客户端信息
func (r *TaobaotaotvvideoplaylistpageAPIRequest) SetSystemInfo(_systemInfo string) error {
	r._systemInfo = _systemInfo
	r.Set("system_info", _systemInfo)
	return nil
}

// GetSystemInfo SystemInfo Getter
func (r TaobaotaotvvideoplaylistpageAPIRequest) GetSystemInfo() string {
	return r._systemInfo
}

// SetPageNo is PageNo Setter
// 当前页（从1开始）
func (r *TaobaotaotvvideoplaylistpageAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaotaotvvideoplaylistpageAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
