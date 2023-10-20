package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemskustatusgetAPIRequest 商品sku上下架查询 API请求
// tmall.item.sku.status.get
//
// 商品sku上下架状态查询
type TmallitemskustatusgetAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// NewTmallitemskustatusgetRequest 初始化TmallitemskustatusgetAPIRequest对象
func NewTmallitemskustatusgetRequest() *TmallitemskustatusgetAPIRequest {
	return &TmallitemskustatusgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemskustatusgetAPIRequest) GetApiMethodName() string {
	return "tmall.item.sku.status.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemskustatusgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemskustatusgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TmallitemskustatusgetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallitemskustatusgetAPIRequest) GetItemId() int64 {
	return r._itemId
}
