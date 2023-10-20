package tmalltrend

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallbranditemuploadAPIRequest 天猫品牌新品同步API API请求
// tmall.brand.item.upload
//
// 支撑天猫品牌将各渠道新品信息同步至平台
type TmallbranditemuploadAPIRequest struct {
	model.Params
	// 需要同步的商品列表
	_itemList []TmallBrandChannelNewItem
}

// NewTmallbranditemuploadRequest 初始化TmallbranditemuploadAPIRequest对象
func NewTmallbranditemuploadRequest() *TmallbranditemuploadAPIRequest {
	return &TmallbranditemuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallbranditemuploadAPIRequest) GetApiMethodName() string {
	return "tmall.brand.item.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallbranditemuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallbranditemuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemList is ItemList Setter
// 需要同步的商品列表
func (r *TmallbranditemuploadAPIRequest) SetItemList(_itemList []TmallBrandChannelNewItem) error {
	r._itemList = _itemList
	r.Set("item_list", _itemList)
	return nil
}

// GetItemList ItemList Getter
func (r TmallbranditemuploadAPIRequest) GetItemList() []TmallBrandChannelNewItem {
	return r._itemList
}
