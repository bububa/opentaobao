package smartstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallpopupstoreitemdiscountpriceAPIRequest 商品优惠价格查询 API请求
// tmall.popupstore.item.discount.price
//
// 商品优惠价格查询
type TmallpopupstoreitemdiscountpriceAPIRequest struct {
	model.Params
	// 商品id列表
	_itemIds []string
}

// NewTmallpopupstoreitemdiscountpriceRequest 初始化TmallpopupstoreitemdiscountpriceAPIRequest对象
func NewTmallpopupstoreitemdiscountpriceRequest() *TmallpopupstoreitemdiscountpriceAPIRequest {
	return &TmallpopupstoreitemdiscountpriceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallpopupstoreitemdiscountpriceAPIRequest) GetApiMethodName() string {
	return "tmall.popupstore.item.discount.price"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallpopupstoreitemdiscountpriceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallpopupstoreitemdiscountpriceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemIds is ItemIds Setter
// 商品id列表
func (r *TmallpopupstoreitemdiscountpriceAPIRequest) SetItemIds(_itemIds []string) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TmallpopupstoreitemdiscountpriceAPIRequest) GetItemIds() []string {
	return r._itemIds
}
