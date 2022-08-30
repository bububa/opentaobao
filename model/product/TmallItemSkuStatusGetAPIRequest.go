package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSkuStatusGetAPIRequest 商品sku上下架查询 API请求
// tmall.item.sku.status.get
//
// 商品sku上下架状态查询
type TmallItemSkuStatusGetAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// NewTmallItemSkuStatusGetRequest 初始化TmallItemSkuStatusGetAPIRequest对象
func NewTmallItemSkuStatusGetRequest() *TmallItemSkuStatusGetAPIRequest {
	return &TmallItemSkuStatusGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemSkuStatusGetAPIRequest) GetApiMethodName() string {
	return "tmall.item.sku.status.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemSkuStatusGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TmallItemSkuStatusGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallItemSkuStatusGetAPIRequest) GetItemId() int64 {
	return r._itemId
}
