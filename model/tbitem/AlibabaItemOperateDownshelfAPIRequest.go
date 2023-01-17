package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItemOperateDownshelfAPIRequest 商品下架 API请求
// alibaba.item.operate.downshelf
//
// 商品下架
type AlibabaItemOperateDownshelfAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// NewAlibabaItemOperateDownshelfRequest 初始化AlibabaItemOperateDownshelfAPIRequest对象
func NewAlibabaItemOperateDownshelfRequest() *AlibabaItemOperateDownshelfAPIRequest {
	return &AlibabaItemOperateDownshelfAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaItemOperateDownshelfAPIRequest) GetApiMethodName() string {
	return "alibaba.item.operate.downshelf"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaItemOperateDownshelfAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaItemOperateDownshelfAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *AlibabaItemOperateDownshelfAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaItemOperateDownshelfAPIRequest) GetItemId() int64 {
	return r._itemId
}
