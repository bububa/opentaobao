package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSkuNewGetAPIRequest 查询sku销售属性标新信息 API请求
// tmall.item.sku.new.get
//
// 查询sku销售属性标新信息
type TmallItemSkuNewGetAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// NewTmallItemSkuNewGetRequest 初始化TmallItemSkuNewGetAPIRequest对象
func NewTmallItemSkuNewGetRequest() *TmallItemSkuNewGetAPIRequest {
	return &TmallItemSkuNewGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemSkuNewGetAPIRequest) GetApiMethodName() string {
	return "tmall.item.sku.new.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemSkuNewGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TmallItemSkuNewGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallItemSkuNewGetAPIRequest) GetItemId() int64 {
	return r._itemId
}
