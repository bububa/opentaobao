package taotv

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaotvCarouselChannelAllAPIRequest 获取所有频道列表 API请求
// taobao.taotv.carousel.channel.all
//
// 获取所有频道列表，按照序号升序
type TaobaoTaotvCarouselChannelAllAPIRequest struct {
	model.Params
	// 系统信息
	_systemInfo string
}

// NewTaobaoTaotvCarouselChannelAllRequest 初始化TaobaoTaotvCarouselChannelAllAPIRequest对象
func NewTaobaoTaotvCarouselChannelAllRequest() *TaobaoTaotvCarouselChannelAllAPIRequest {
	return &TaobaoTaotvCarouselChannelAllAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTaotvCarouselChannelAllAPIRequest) Reset() {
	r._systemInfo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTaotvCarouselChannelAllAPIRequest) GetApiMethodName() string {
	return "taobao.taotv.carousel.channel.all"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTaotvCarouselChannelAllAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTaotvCarouselChannelAllAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSystemInfo is SystemInfo Setter
// 系统信息
func (r *TaobaoTaotvCarouselChannelAllAPIRequest) SetSystemInfo(_systemInfo string) error {
	r._systemInfo = _systemInfo
	r.Set("system_info", _systemInfo)
	return nil
}

// GetSystemInfo SystemInfo Getter
func (r TaobaoTaotvCarouselChannelAllAPIRequest) GetSystemInfo() string {
	return r._systemInfo
}

var poolTaobaoTaotvCarouselChannelAllAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTaotvCarouselChannelAllRequest()
	},
}

// GetTaobaoTaotvCarouselChannelAllRequest 从 sync.Pool 获取 TaobaoTaotvCarouselChannelAllAPIRequest
func GetTaobaoTaotvCarouselChannelAllAPIRequest() *TaobaoTaotvCarouselChannelAllAPIRequest {
	return poolTaobaoTaotvCarouselChannelAllAPIRequest.Get().(*TaobaoTaotvCarouselChannelAllAPIRequest)
}

// ReleaseTaobaoTaotvCarouselChannelAllAPIRequest 将 TaobaoTaotvCarouselChannelAllAPIRequest 放入 sync.Pool
func ReleaseTaobaoTaotvCarouselChannelAllAPIRequest(v *TaobaoTaotvCarouselChannelAllAPIRequest) {
	v.Reset()
	poolTaobaoTaotvCarouselChannelAllAPIRequest.Put(v)
}
