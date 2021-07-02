package tmalltrend

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallBrandItemUploadAPIRequest 天猫品牌新品同步API API请求
// tmall.brand.item.upload
//
// 支撑天猫品牌将各渠道新品信息同步至平台
type TmallBrandItemUploadAPIRequest struct {
	model.Params
	// 需要同步的商品列表
	_itemList []TmallBrandChannelNewItem
}

// NewTmallBrandItemUploadRequest 初始化TmallBrandItemUploadAPIRequest对象
func NewTmallBrandItemUploadRequest() *TmallBrandItemUploadAPIRequest {
	return &TmallBrandItemUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallBrandItemUploadAPIRequest) GetApiMethodName() string {
	return "tmall.brand.item.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallBrandItemUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetItemList is ItemList Setter
// 需要同步的商品列表
func (r *TmallBrandItemUploadAPIRequest) SetItemList(_itemList []TmallBrandChannelNewItem) error {
	r._itemList = _itemList
	r.Set("item_list", _itemList)
	return nil
}

// GetItemList ItemList Getter
func (r TmallBrandItemUploadAPIRequest) GetItemList() []TmallBrandChannelNewItem {
	return r._itemList
}
