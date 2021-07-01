package tmalltrend

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallBrandItemUploadAPIRequest
天猫品牌新品同步API API请求
tmall.brand.item.upload

支撑天猫品牌将各渠道新品信息同步至平台 */
type TmallBrandItemUploadAPIRequest struct {
	model.Params
	// 需要同步的商品列表
	_itemList []TmallBrandChannelNewItem
}

// New
