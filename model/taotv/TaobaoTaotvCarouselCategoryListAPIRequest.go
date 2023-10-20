package taotv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaotvCarouselCategoryListAPIRequest 获取轮播分类列表 API请求
// taobao.taotv.carousel.category.list
//
// 获取轮播分类列表
type TaobaoTaotvCarouselCategoryListAPIRequest struct {
	model.Params
	// 设备信息
	_systemInfo string
}

// NewTaobaoTaotvCarouselCategoryListRequest 初始化TaobaoTaotvCarouselCategoryListAPIRequest对象
func NewTaobaoTaotvCarouselCategoryListRequest() *TaobaoTaotvCarouselCategoryListAPIRequest {
	return &TaobaoTaotvCarouselCategoryListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTaotvCarouselCategoryListAPIRequest) GetApiMethodName() string {
	return "taobao.taotv.carousel.category.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTaotvCarouselCategoryListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTaotvCarouselCategoryListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSystemInfo is SystemInfo Setter
// 设备信息
func (r *TaobaoTaotvCarouselCategoryListAPIRequest) SetSystemInfo(_systemInfo string) error {
	r._systemInfo = _systemInfo
	r.Set("system_info", _systemInfo)
	return nil
}

// GetSystemInfo SystemInfo Getter
func (r TaobaoTaotvCarouselCategoryListAPIRequest) GetSystemInfo() string {
	return r._systemInfo
}
