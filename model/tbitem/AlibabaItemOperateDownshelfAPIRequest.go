package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaitemoperatedownshelfAPIRequest 商品下架 API请求
// alibaba.item.operate.downshelf
//
// 商品下架
type AlibabaitemoperatedownshelfAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// NewAlibabaitemoperatedownshelfRequest 初始化AlibabaitemoperatedownshelfAPIRequest对象
func NewAlibabaitemoperatedownshelfRequest() *AlibabaitemoperatedownshelfAPIRequest {
	return &AlibabaitemoperatedownshelfAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaitemoperatedownshelfAPIRequest) GetApiMethodName() string {
	return "alibaba.item.operate.downshelf"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaitemoperatedownshelfAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaitemoperatedownshelfAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *AlibabaitemoperatedownshelfAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaitemoperatedownshelfAPIRequest) GetItemId() int64 {
	return r._itemId
}
