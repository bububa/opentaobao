package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemupdatesimpleschemagetAPIRequest 官网同购编辑商品的get接口 API请求
// tmall.item.update.simpleschema.get
//
// 官网同购编辑商品的get接口
type TmallitemupdatesimpleschemagetAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// NewTmallitemupdatesimpleschemagetRequest 初始化TmallitemupdatesimpleschemagetAPIRequest对象
func NewTmallitemupdatesimpleschemagetRequest() *TmallitemupdatesimpleschemagetAPIRequest {
	return &TmallitemupdatesimpleschemagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemupdatesimpleschemagetAPIRequest) GetApiMethodName() string {
	return "tmall.item.update.simpleschema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemupdatesimpleschemagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemupdatesimpleschemagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品id
func (r *TmallitemupdatesimpleschemagetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallitemupdatesimpleschemagetAPIRequest) GetItemId() int64 {
	return r._itemId
}
