package taotv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTaotvCarouselCategoryListAPIRequest
获取轮播分类列表 API请求
taobao.taotv.carousel.category.list

获取轮播分类列表 */
type TaobaoTaotvCarouselCategoryListAPIRequest struct {
	model.Params
	// 设备信息
	_systemInfo string
}

// New
