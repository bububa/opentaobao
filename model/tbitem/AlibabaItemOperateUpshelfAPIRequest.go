package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaitemoperateupshelfAPIRequest 商品上架 API请求
// alibaba.item.operate.upshelf
//
// 商品上架
type AlibabaitemoperateupshelfAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
	// 商品库存
	_quantity int64
}

// NewAlibabaitemoperateupshelfRequest 初始化AlibabaitemoperateupshelfAPIRequest对象
func NewAlibabaitemoperateupshelfRequest() *AlibabaitemoperateupshelfAPIRequest {
	return &AlibabaitemoperateupshelfAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaitemoperateupshelfAPIRequest) GetApiMethodName() string {
	return "alibaba.item.operate.upshelf"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaitemoperateupshelfAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaitemoperateupshelfAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *AlibabaitemoperateupshelfAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaitemoperateupshelfAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetQuantity is Quantity Setter
// 商品库存
func (r *AlibabaitemoperateupshelfAPIRequest) SetQuantity(_quantity int64) error {
	r._quantity = _quantity
	r.Set("quantity", _quantity)
	return nil
}

// GetQuantity Quantity Getter
func (r AlibabaitemoperateupshelfAPIRequest) GetQuantity() int64 {
	return r._quantity
}
