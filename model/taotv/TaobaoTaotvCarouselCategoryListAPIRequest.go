package taotv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotaotvcarouselcategorylistAPIRequest 获取轮播分类列表 API请求
// taobao.taotv.carousel.category.list
//
// 获取轮播分类列表
type TaobaotaotvcarouselcategorylistAPIRequest struct {
	model.Params
	// 设备信息
	_systemInfo string
}

// NewTaobaotaotvcarouselcategorylistRequest 初始化TaobaotaotvcarouselcategorylistAPIRequest对象
func NewTaobaotaotvcarouselcategorylistRequest() *TaobaotaotvcarouselcategorylistAPIRequest {
	return &TaobaotaotvcarouselcategorylistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotaotvcarouselcategorylistAPIRequest) GetApiMethodName() string {
	return "taobao.taotv.carousel.category.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotaotvcarouselcategorylistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotaotvcarouselcategorylistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSystemInfo is SystemInfo Setter
// 设备信息
func (r *TaobaotaotvcarouselcategorylistAPIRequest) SetSystemInfo(_systemInfo string) error {
	r._systemInfo = _systemInfo
	r.Set("system_info", _systemInfo)
	return nil
}

// GetSystemInfo SystemInfo Getter
func (r TaobaotaotvcarouselcategorylistAPIRequest) GetSystemInfo() string {
	return r._systemInfo
}
