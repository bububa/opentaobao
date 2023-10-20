package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItemOperateUpshelfAPIRequest 商品上架 API请求
// alibaba.item.operate.upshelf
//
// 商品上架
type AlibabaItemOperateUpshelfAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
	// 商品库存
	_quantity int64
}

// NewAlibabaItemOperateUpshelfRequest 初始化AlibabaItemOperateUpshelfAPIRequest对象
func NewAlibabaItemOperateUpshelfRequest() *AlibabaItemOperateUpshelfAPIRequest {
	return &AlibabaItemOperateUpshelfAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaItemOperateUpshelfAPIRequest) Reset() {
	r._itemId = 0
	r._quantity = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaItemOperateUpshelfAPIRequest) GetApiMethodName() string {
	return "alibaba.item.operate.upshelf"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaItemOperateUpshelfAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaItemOperateUpshelfAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *AlibabaItemOperateUpshelfAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaItemOperateUpshelfAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetQuantity is Quantity Setter
// 商品库存
func (r *AlibabaItemOperateUpshelfAPIRequest) SetQuantity(_quantity int64) error {
	r._quantity = _quantity
	r.Set("quantity", _quantity)
	return nil
}

// GetQuantity Quantity Getter
func (r AlibabaItemOperateUpshelfAPIRequest) GetQuantity() int64 {
	return r._quantity
}

var poolAlibabaItemOperateUpshelfAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaItemOperateUpshelfRequest()
	},
}

// GetAlibabaItemOperateUpshelfRequest 从 sync.Pool 获取 AlibabaItemOperateUpshelfAPIRequest
func GetAlibabaItemOperateUpshelfAPIRequest() *AlibabaItemOperateUpshelfAPIRequest {
	return poolAlibabaItemOperateUpshelfAPIRequest.Get().(*AlibabaItemOperateUpshelfAPIRequest)
}

// ReleaseAlibabaItemOperateUpshelfAPIRequest 将 AlibabaItemOperateUpshelfAPIRequest 放入 sync.Pool
func ReleaseAlibabaItemOperateUpshelfAPIRequest(v *AlibabaItemOperateUpshelfAPIRequest) {
	v.Reset()
	poolAlibabaItemOperateUpshelfAPIRequest.Put(v)
}
