package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemskunewgetAPIRequest 查询sku销售属性标新信息 API请求
// tmall.item.sku.new.get
//
// 查询sku销售属性标新信息
type TmallitemskunewgetAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// NewTmallitemskunewgetRequest 初始化TmallitemskunewgetAPIRequest对象
func NewTmallitemskunewgetRequest() *TmallitemskunewgetAPIRequest {
	return &TmallitemskunewgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemskunewgetAPIRequest) GetApiMethodName() string {
	return "tmall.item.sku.new.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemskunewgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemskunewgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TmallitemskunewgetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallitemskunewgetAPIRequest) GetItemId() int64 {
	return r._itemId
}
